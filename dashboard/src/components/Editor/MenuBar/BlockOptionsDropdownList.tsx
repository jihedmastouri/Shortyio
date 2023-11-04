import { useCallback } from "react";
import Select from "@ui/Select";

import {
  ActivityLogIcon,
  CodeIcon,
  HeadingIcon,
  ListBulletIcon,
  QuoteIcon,
  TextIcon,
} from "@radix-ui/react-icons";

import {
  $createParagraphNode,
  $getSelection,
  $isRangeSelection,
} from "lexical";
import { $wrapNodes } from "@lexical/selection";
import {
  INSERT_ORDERED_LIST_COMMAND,
  INSERT_UNORDERED_LIST_COMMAND,
  REMOVE_LIST_COMMAND,
} from "@lexical/list";
import { $createHeadingNode, $createQuoteNode } from "@lexical/rich-text";
import { $createCodeNode } from "@lexical/code";

type Props = {
  editor: any;
  blockType: string;
};

export default function BlockOptionsDropdownList({ editor, blockType }: Props) {
  const textFormatChanged = useCallback(
    (newValue: string) => {
      if (newValue == blockType) {
        if (newValue == "ul" || newValue == "ol")
          editor.dispatchCommand(REMOVE_LIST_COMMAND);
        return;
      }

      if (newValue == "ul") {
        editor.dispatchCommand(INSERT_UNORDERED_LIST_COMMAND);
        return;
      }

      if (newValue == "ol") {
        editor.dispatchCommand(INSERT_ORDERED_LIST_COMMAND);
        return;
      }

      editor.update(() => {
        const selection = $getSelection();

        if ($isRangeSelection(selection)) {
          switch (newValue) {
            case "paragraph":
              $wrapNodes(selection, () => $createParagraphNode());
              break;
            case "h2":
            case "h3":
            case "h4":
              $wrapNodes(selection, () => $createHeadingNode(newValue));
              break;
            case "quote":
              $wrapNodes(selection, () => $createQuoteNode());
              break;
            case "code":
              $wrapNodes(selection, () => $createCodeNode());
              break;
            default:
              $wrapNodes(selection, () => $createParagraphNode());
              break;
          }
        }
      });
    },
    [blockType, editor]
  );

  return (
    <Select
      ariaLabel="Formatting Options"
      selected={blockType}
      onChange={textFormatChanged}
      items={[
        {
          name: "paragraph",
          content: (
            <span className="flex items-center  justify-between">
              <TextIcon className="w-4" />
              <span className="ml-2">Normal</span>
            </span>
          ),
        },
        {
          name: "h2",
          content: (
            <span className="flex items-center  justify-between">
              <HeadingIcon className="w-4" />
              <span className="ml-2">Large Heading</span>
            </span>
          ),
        },
        {
          name: "h3",
          content: (
            <span className="flex items-center  justify-between">
              <HeadingIcon className="w-4" />
              <span className="ml-2">Medium Heading</span>
            </span>
          ),
        },
        {
          name: "h4",
          content: (
            <span className="flex items-center  justify-between">
              <HeadingIcon className="w-4" />
              <span className="ml-2">Small Heading</span>
            </span>
          ),
        },
        {
          name: "h5",
          content: (
            <span className="flex items-center  justify-between">
              <HeadingIcon className="w-4" />
              <span className="ml-2">Smaller Heading</span>
            </span>
          ),
        },
        {
          name: "ul",
          content: (
            <span className="flex items-center  justify-between">
              <ListBulletIcon className="w-4" />
              <span className="ml-2">Bullet List</span>
            </span>
          ),
        },
        {
          name: "ol",
          content: (
            <span className="flex items-center  justify-between">
              <ActivityLogIcon className="w-4" />
              <span className="ml-2">Numbered List</span>
            </span>
          ),
        },
        {
          name: "quote",
          content: (
            <span className="flex items-center  justify-between">
              <QuoteIcon className="w-4" />
              <span className="ml-2">Quote</span>
            </span>
          ),
        },
        {
          name: "code",
          content: (
            <span className="flex items-center  justify-between">
              <CodeIcon className="w-4" />
              <span className="ml-2">Code Block</span>
            </span>
          ),
        },
      ]}
    ></Select>
  );
}
