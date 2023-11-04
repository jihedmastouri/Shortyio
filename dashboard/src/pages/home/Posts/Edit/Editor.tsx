import AppEditor from "@Editor";
import {
  ExternalLinkIcon,
  Pencil1Icon,
  Pencil2Icon,
  PlusIcon,
  TrashIcon,
} from "@radix-ui/react-icons";
import * as Toolbar from "@radix-ui/react-toolbar";
import AppToolTips from "@ui/Tooltip";
import Select from "@ui/Select";
import Island from "pages/home/components/Island";
import { useId, useState } from "react";
import { useBlockContent, useBlockMeta } from "./utils";
import Loading from "components/Loading";
import Error from "components/Error";

type TextualType = {
  id: string;
  name: string;
  text: string;
  hint: string;
};

// TODO: Add image Elements
type ImageType = {
  id: string;
  alt: string;
  file: string;
};

type Props = {
  id: string;
};

const languages = [
  { content: "English", name: "en_US" },
  { content: "Spanish", name: "es_SP" },
  { content: "french", name: "fr_FR" },
  { content: "french Canda", name: "fr_CA" },
  { content: "Arabic", name: "ar_AR" },
  { content: "Tunisian", name: "ar_TN" },
];

export default function Editor({ id }: Props) {
  const [lang, setLang] = useState(languages[0].name);

  const { data, isLoading, error } = useBlockMeta(id, "en_US");
  const {
    data: content,
    isLoading: isLoadingContent,
    error: errorContent,
  } = useBlockContent(id, lang);

  if (isLoading) return <Loading />;
  if (error || !data) return <Error />;

  const post = {
    title: "this is a big title a reallt big one soooo big",
    lang: "language",
    langCode: "en",
  };

  const TextElement = {
    id: "1",
    name: "This is a Title",
    text: "This is a text",
    hint: "You lost?",
  };

  const newElement = () => {
    alert("new post");
  };

  return (
    <>
      <h1 className="font-semibold text-4xl max-w-[600px] mb-2 py-2 mx-auto text-center text-gray-800 italic capitalize">
        "{data.name}."
      </h1>
      <Toolbar.Root className="flex justify-start items-center bg-white rounded-full shadow-lg px-5 py-2 mx-auto my-2 w-11/12 sticky top-2">
        <Toolbar.Button className="transition h-5 w-5 rounded-full border border-gray-500">
          <AppToolTips content="Create a new language post">
            <PlusIcon />
          </AppToolTips>
        </Toolbar.Button>
        <Toolbar.Separator className="w-px h-5 bg-gray-300 mx-2" />
        <h2 className="text-s-grass font-bold">Language:</h2>
        <Select
          items={languages}
          onChange={(e) => setLang(e)}
          selected={lang}
          ariaLabel="Select a language"
          className="shadow rounded-md border mx-2 px-2 py-1 text-sm font-medium text-gray-700 hover:bg-gray-100"
        />
        <Toolbar.Button className="ml-auto transition rounded-md bg-s-gopher text-white mr-2 px-1 hover:shadow select-none font-bold">
          Save
        </Toolbar.Button>
      </Toolbar.Root>
      {isLoadingContent ? (
        <Loading />
      ) : errorContent ? (
        <Error />
      ) : content?.content ? (
        <>
          <div className="w-8/12 mx-auto">
            {content.content.map((el) => {
              if (el["text"]) {
                const textType = el.text.type;
                if (textType === "TEXTUAL_TYPE_TEXT") {
                  return <ElementText {...TextElement} />;
                } else if (textType === "TEXTUAL_TYPE_RICH") {
                  return <ElementRich {...TextElement} />;
                }
              }
              return null;
            })}
          </div>
          <div
            onClick={newElement}
            className="flex rounded-full w-10 h-10 p-1 m-1 border items-center hover:text-s-gopher hover:border-s-gopher mx-auto mb-10"
          >
            <PlusIcon />
          </div>
        </>
      ) : (
        <EmptyPost handler={newElement} />
      )}
    </>
  );
}

function EmptyPost(props: { handler: () => void }) {
  return (
    <div
      onClick={props.handler}
      className="flex flex-col justify-center items-center h-full mx-auto my-10 cursor-pointer py-2 px-10 border border-dashed border-gray-700 text-gray-500 hover:text-s-gopher hover:border-s-gopher group"
    >
      <p className="text-3xl font-semibold text-gray-700 leading-loose group-hover:text-s-gopher">
        Such Empty. Much Sad.
      </p>
      <p className="text-gray-500 group-hover:text-s-gopher">
        Click Here To Start Editing You Post
      </p>
      <div className="flex rounded-full w-7 h-7 p-1 m-1 border items-center">
        <PlusIcon />
      </div>
    </div>
  );
}

function NewElementSelector() {
  return (
    <Island className="w-1/2 mx-auto flex flex-col">
      <h2 className="font-bold mx-auto my-1">Choose element type:</h2>
      <div className="flex justify-evenly items-center">
        <div
          onClick={() => alert("new text")}
          className="mx-5 w-250 flex justify-center items-center cursor-pointer py-2 px-10 rounded-md border border-dotted border-gray-700 text-gray-500 hover:text-s-gopher hover:border-s-gopher group"
        >
          <span className="w-7 h-7">
            <Pencil1Icon />
          </span>
          <p className="font-semibold text-gray-700 leading-loose group-hover:text-s-gopher">
            Simple Text Element
          </p>
        </div>
        <div
          onClick={() => alert("new text")}
          className="mx-5 w-250 flex justify-center items-center cursor-pointer py-2 px-10 rounded-md border border-dotted border-gray-700 text-gray-500 hover:text-s-gopher hover:border-s-gopher group"
        >
          <span className="w-7 h-7">
            <Pencil2Icon />
          </span>
          <p className="font-semibold text-gray-700 leading-loose group-hover:text-s-gopher">
            Rich Text Element
          </p>
        </div>
      </div>
    </Island>
  );
}

function ElementText(props: TextualType) {
  const [textName, setTextName] = useState(props.name);
  const [textContent, setTextContent] = useState(props.text);

  const inputNameId = useId();
  const inputContentId = useId();
  return (
    <Island title={props.hint} className="p-2 mb-2">
      <div className="flex justify-between items-center">
        <div className="w-10/12 flex items-center mt-2">
          <label className="font-bold select-none" htmlFor={inputNameId}>
            Name:
          </label>
          <input
            className="h-full w-full mx-2 shadow rounded-md border px-2 py-1 font-medium outline-s-gopher"
            type="text"
            value={textName}
            id={inputNameId}
            onChange={(e) => setTextName(e.target.value)}
          />
        </div>
        <span className="w-5 h-5 hover:text-red-500">
          <TrashIcon />
        </span>
      </div>
      <div className="h-px w-full my-3 mx-auto bg-gray-100 shadow"></div>
      <div className="flex items-center mb-4 mt-2">
        <label className="font-bold select-none" htmlFor={inputContentId}>
          Content:
        </label>
        <input
          className="h-full w-full mx-2 shadow rounded-md border px-2 py-1 font-medium outline-s-gopher"
          type="text"
          value={textContent}
          id={inputContentId}
          onChange={(e) => setTextContent(e.target.value)}
        />
      </div>
    </Island>
  );
}

function ElementRich(props: TextualType) {
  const [textName, setTextName] = useState(props.name);
  const [textContent, setTextContent] = useState(props.text);

  const inputNameId = useId();
  return (
    <Island title={props.hint} className="p-2">
      <div className="flex justify-between items-center">
        <div className="w-10/12 flex items-center mt-2">
          <label className="font-bold select-none" htmlFor={inputNameId}>
            Name:
          </label>
          <input
            className="h-full w-full mx-2 shadow rounded-md border px-2 py-1 font-medium outline-s-gopher"
            type="text"
            value={textName}
            id={inputNameId}
            onChange={(e) => setTextName(e.target.value)}
          />
        </div>
        <div className="flex">
          <span className="w-5 h-5 hover:text-red-500">
            <TrashIcon />
          </span>
          <span className="w-5 h-5 hover:text-sky-500">
            <ExternalLinkIcon />
          </span>
        </div>
      </div>
      <div className="h-px w-full my-3 mx-auto bg-gray-100 shadow"></div>
      <AppEditor />
    </Island>
  );
}
