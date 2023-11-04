import { useLexicalComposerContext } from "@lexical/react/LexicalComposerContext";
import * as Toolbar from "@radix-ui/react-toolbar";
import { useCallback, useEffect, useMemo, useRef, useState } from "react";

import {
  CAN_REDO_COMMAND,
  CAN_UNDO_COMMAND,
  REDO_COMMAND,
  UNDO_COMMAND,
  SELECTION_CHANGE_COMMAND,
  FORMAT_TEXT_COMMAND,
  FORMAT_ELEMENT_COMMAND,
  $getSelection,
  $isRangeSelection,
  $getNodeByKey,
} from "lexical";

import { $isLinkNode, TOGGLE_LINK_COMMAND } from "@lexical/link";

import { $isParentElementRTL, $isAtNodeEnd } from "@lexical/selection";

import { $getNearestNodeOfType, mergeRegister } from "@lexical/utils";

import { $isListNode, ListNode } from "@lexical/list";

import { createPortal } from "react-dom";
import { $isHeadingNode } from "@lexical/rich-text";

import {
  $isCodeNode,
  getDefaultCodeLanguage,
  getCodeLanguages,
} from "@lexical/code";

import {
  CodeIcon,
  FontBoldIcon,
  FontItalicIcon,
  Link1Icon,
  ResetIcon,
  StrikethroughIcon,
  TextAlignCenterIcon,
  TextAlignJustifyIcon,
  TextAlignLeftIcon,
  TextAlignRightIcon,
  UnderlineIcon,
} from "@radix-ui/react-icons";
import Select from "@ui/Select";
import BlockOptionsDropdownList from "./BlockOptionsDropdownList";
import FloatingLinkEditor from "./FloatingEditor";
import { getSelectedNode } from "./utils";
import * as Separator from "@radix-ui/react-separator";

const LowPriority = 1;

const supportedBlockTypes = new Set([
  "paragraph",
  "quote",
  "code",
  "h2",
  "h3",
  "h4",
  "ul",
  "ol",
]);

type formatType = "bold" | "italic" | "strikethrough" | "underline";
// type alignType = "left" | "center" | "right" | "justify"

export default function ToolbarPlugin({ className }: { className: string }) {
  const [editor] = useLexicalComposerContext();
  const toolbarRef = useRef(null);

  const [canUndo, setCanUndo] = useState(false);
  const [canRedo, setCanRedo] = useState(false);

  const [blockType, setBlockType] = useState("paragraph");

  const [selectedElementKey, setSelectedElementKey] = useState(null);

  // these are of type `formatType` and `alignType`. But I dont't know how to cast it to string[].
  const [textFormat, setTextFormat] = useState<string[]>([]);
  const [textAlign, setTextAlign] = useState<string>("left");

  const [isCode, setIsCode] = useState(false);
  const [codeLanguage, setCodeLanguage] = useState("");
  const [isRTL, setIsRTL] = useState(false);
  const [isLink, setIsLink] = useState(false);

  const updateToolbar = useCallback(() => {
    const selection = $getSelection();

    if ($isRangeSelection(selection)) {
      const anchorNode = selection.anchor.getNode();

      const element =
        anchorNode.getKey() === "root"
          ? anchorNode
          : anchorNode.getTopLevelElementOrThrow();
      const elementKey = element.getKey();
      const elementDOM = editor.getElementByKey(elementKey);

      if (elementDOM !== null) {
        setSelectedElementKey(elementKey);

        if ($isListNode(element)) {
          const parentList = $getNearestNodeOfType(anchorNode, ListNode);
          const type = parentList ? parentList.getTag() : element.getTag();
          setBlockType(type);
        } else {
          const type = $isHeadingNode(element)
            ? element.getTag()
            : element.getType();
          setBlockType(type);
          if ($isCodeNode(element)) {
            setCodeLanguage(element.getLanguage() || getDefaultCodeLanguage());
          }
        }
      }

      const format: formatType[] = [];
      if (selection.hasFormat("bold")) format.push("bold");
      if (selection.hasFormat("italic")) format.push("italic");
      if (selection.hasFormat("underline")) format.push("underline");
      if (selection.hasFormat("strikethrough")) format.push("strikethrough");
      setTextFormat(format);

      setIsCode(selection.hasFormat("code"));
      setIsRTL($isParentElementRTL(selection));

      // Update links
      const node = getSelectedNode(selection);
      const parent = node.getParent();
      if ($isLinkNode(parent) || $isLinkNode(node)) {
        setIsLink(true);
      } else {
        setIsLink(false);
      }
    }
  }, [editor]);

  const textFormatChanged = (val: string[]) => {
    const difference = val.filter((x) => !textFormat.includes(x));
    difference.forEach((el) => {
      editor.dispatchCommand(FORMAT_TEXT_COMMAND, el);
    });
  };

  const textAlignChanged = (val: string) => {
    setTextAlign(val);
    editor.dispatchCommand(FORMAT_ELEMENT_COMMAND, val);
  };

  useEffect(() => {
    return mergeRegister(
      editor.registerUpdateListener(({ editorState }) => {
        editorState.read(() => {
          updateToolbar();
        });
      }),
      editor.registerCommand(
        SELECTION_CHANGE_COMMAND,
        (_payload, newEditor) => {
          updateToolbar();
          return false;
        },
        LowPriority
      ),
      editor.registerCommand(
        CAN_UNDO_COMMAND,
        (payload) => {
          setCanUndo(payload);
          return false;
        },
        LowPriority
      ),
      editor.registerCommand(
        CAN_REDO_COMMAND,
        (payload) => {
          setCanRedo(payload);
          return false;
        },
        LowPriority
      )
    );
  }, [editor, updateToolbar]);

  const codeLanguges = useMemo(
    () => getCodeLanguages().map((lang) => ({ name: lang, content: lang })),
    []
  );

  const onCodeLanguageSelect = useCallback(
    (e) => {
      editor.update(() => {
        if (selectedElementKey !== null) {
          const node = $getNodeByKey(selectedElementKey);
          if ($isCodeNode(node)) {
            node.setLanguage(e.target.value);
          }
        }
      });
    },
    [editor, selectedElementKey]
  );

  const insertLink = useCallback(() => {
    if (!isLink) {
      editor.dispatchCommand(TOGGLE_LINK_COMMAND, "https://");
    } else {
      editor.dispatchCommand(TOGGLE_LINK_COMMAND, null);
    }
  }, [editor, isLink]);

  return (
    <Toolbar.Root
      ref={toolbarRef}
      className={className}
      aria-label="Formatting options"
    >
      <Toolbar.Button
        disabled={!canUndo}
        onClick={() => {
          editor.dispatchCommand(UNDO_COMMAND);
        }}
        className={canUndo ? "fill-gray-50" : ""}
        aria-label="Undo"
      >
        <ResetIcon />
      </Toolbar.Button>
      <Toolbar.Button
        disabled={!canRedo}
        onClick={() => {
          editor.dispatchCommand(REDO_COMMAND);
        }}
        className={canRedo ? "text-gray-500" : ""}
        aria-label="Redo"
      >
        <ResetIcon className="scale-x-[-1]" />
      </Toolbar.Button>

      <Toolbar.Separator className="w-px h-6 bg-gray-300" />

      {supportedBlockTypes.has(blockType) && (
        <span className="flex justify-between">
          <BlockOptionsDropdownList editor={editor} blockType={blockType} />
        </span>
      )}
      {blockType === "code" ? (
        <span className="flex justify-between">
          <Select
            ariaLabel="Programing Languages"
            onChange={() => onCodeLanguageSelect}
            items={codeLanguges}
            selected={codeLanguage}
          />
        </span>
      ) : (
        <>
          <Toolbar.Separator className="w-px h-6 bg-gray-300" />
          <Toolbar.ToggleGroup
            value={textFormat}
            onValueChange={textFormatChanged}
            type="multiple"
            aria-label="Text formatting"
          >
            <Toolbar.ToggleItem
              value="bold"
              className={
                "toolbar-item spaced " +
                (textFormat.includes("bold") ? "active" : "")
              }
              aria-label="Format Bold"
            >
              <FontBoldIcon />
            </Toolbar.ToggleItem>
            <Toolbar.ToggleItem
              value="italic"
              className={
                "toolbar-item spaced " +
                (textFormat.includes("italic") ? "active" : "")
              }
              aria-label="Format Italics"
            >
              <FontItalicIcon />
            </Toolbar.ToggleItem>
            <Toolbar.ToggleItem
              className={
                "toolbar-item spaced " +
                (textFormat.includes("strikethrough") ? "active" : "")
              }
              value="strikethrough"
              aria-label="Strike through"
            >
              <StrikethroughIcon />
            </Toolbar.ToggleItem>
            <Toolbar.ToggleItem
              className={
                "toolbar-item spaced " +
                (textFormat.includes("underline") ? "active" : "")
              }
              value="underline"
              a-label="Format Underline"
            >
              <UnderlineIcon />
            </Toolbar.ToggleItem>
          </Toolbar.ToggleGroup>

          <Toolbar.Separator className="w-px h-6 mx-2 bg-gray-300" />
          <Toolbar.Button
            onClick={() => {
              editor.dispatchCommand(FORMAT_TEXT_COMMAND, "code");
            }}
            className={"toolbar-item spaced " + (isCode ? "active" : "")}
            aria-label="Insert Code"
          >
            <CodeIcon />
          </Toolbar.Button>

          <Toolbar.Button
            onClick={insertLink}
            className={"toolbar-item spaced " + (isLink ? "active" : "")}
            aria-label="Insert Link"
          >
            <Link1Icon />
          </Toolbar.Button>

          {isLink &&
            createPortal(<FloatingLinkEditor editor={editor} />, document.body)}

          <Toolbar.Separator className="w-px h-6 mx-2 bg-gray-300" />

          <Toolbar.ToggleGroup
            value={textAlign}
            onValueChange={textAlignChanged}
            type="single"
            aria-label="Text formatting"
          >
            <Toolbar.ToggleItem
              value="left"
              className="toolbar-item spaced"
              aria-label="Left Align"
            >
              <TextAlignLeftIcon />
            </Toolbar.ToggleItem>
            <Toolbar.ToggleItem
              value="center"
              className="toolbar-item spaced"
              aria-label="Center Align"
            >
              <TextAlignCenterIcon />
            </Toolbar.ToggleItem>
            <Toolbar.ToggleItem
              value="right"
              className="toolbar-item spaced"
              aria-label="Right Align"
            >
              <TextAlignRightIcon />
            </Toolbar.ToggleItem>
            <Toolbar.ToggleItem
              value="justify"
              className="toolbar-item"
              aria-label="Justify Align"
            >
              <TextAlignJustifyIcon />
            </Toolbar.ToggleItem>
          </Toolbar.ToggleGroup>
        </>
      )}
    </Toolbar.Root>
  );
}
