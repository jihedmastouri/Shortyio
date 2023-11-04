import {
  GearIcon,
  Pencil1Icon,
  PersonIcon,
  UploadIcon,
} from "@radix-ui/react-icons";
import Island from "pages/home/components/Island";
import { FormEvent, useEffect, useRef, useState } from "react";
import InputElm from "./InputElm";
import SwitchToggle from "@ui/Switch";
import Select from "@ui/Select";
import { Link } from "react-router-dom";
import { useMutation } from "@tanstack/react-query";
import axios from "axios";
import MultiSelect, { GroupBase, Options } from "react-select";
import { useBlockMeta, useBlockRules } from "./utils";
import Loading from "components/Loading";
import Error from "components/Error";

export default function Settings({ id }: { id: string }) {
  const parentRef = useRef<HTMLDivElement>(null);

  return (
    <>
      <div className="outline-none flex flex-col" ref={parentRef}>
        <div className="grid grid-cols-3 gap-1">
          <MetaSection id={id} />
          <SideSettings id={id} />
        </div>
      </div>
    </>
  );
}

function MetaSection({ id }: { id: string }) {
  const options: Options<GroupBase<string>> = [
    { label: "Personal", value: "Personal" },
    { label: "Business", value: "Business" },
  ];

  const { data, isLoading, error } = useBlockMeta(id);
  const [form, setForm] = useState({ ...data });

  useEffect(() => {
    let timer = setTimeout(() => {
      saveInput();
    }, 700);
    return () => clearTimeout(timer);
  }, [form]);

  const saveInput = () => {
    console.log("Saving data", form);
    axios.put(`/cmd/block/${id}`, { ...form }).then((res) => {
      console.log(res);
    });
  };

  const handleFormChange = (e: any) => {
    const { name, value } = e.currentTarget;
    setForm((prev) => ({ ...prev, [name]: value }));
  };

  const handleCategChange = (e: any) => {
    console.log(e);
    axios.post(`/cmd/block/${id}/categ/${e.value}`);
    setForm((prev) => ({ ...prev, categories: e.map((e: any) => e.value) }));
  };

  const handleTagsChange = (e: any) => {
    console.log(e);
    axios.post(`/cmd/block/${id}/tag/${e.value}`);
    setForm((prev) => ({ ...prev, tags: e.map((e: any) => e.value) }));
  };

  if (isLoading) return <Loading />;

  if (error || !data) return <Error />;

  return (
    <div className="col-span-2 flex flex-col items-center w-full">
      <section
        aria-roledescription="Meta infos"
        className="mb-5 flex flex-col items-center w-full"
      >
        <h2 className="font-bold text-xl self-start my-2">Meta Infos</h2>
        <form method="POST" className="w-11/12">
          <InputElm name="name" id="title">
            <input
              className="p-1 rounded shadow text-black w-3/4 h-10"
              type="text"
              id="title"
              name="name"
              placeholder="Page Title"
              defaultValue={form.name}
              onChange={handleFormChange}
            />
          </InputElm>
          <InputElm name="Description" id="desc">
            <textarea
              className="p-1 rounded shadow text-black w-3/4 h-20"
              id="desc"
              name="description"
              placeholder="Short Description (150 words max)"
              onChange={handleFormChange}
              defaultValue={form.description}
            />
          </InputElm>
          <InputElm name="Tags" id="tags">
            <span className="p-1 rounded shadow text-black w-3/4">
              <MultiSelect
                defaultValue={form.tags}
                onChange={handleTagsChange}
                name="tags"
                isMulti
                options={options}
              ></MultiSelect>
            </span>
          </InputElm>
          <InputElm name="Categories" id="categs">
            <span className="p-1 rounded shadow text-black w-3/4">
              <MultiSelect
                defaultValue={form.categories}
                onChange={handleCategChange}
                name="categories"
                isMulti
                options={options}
              ></MultiSelect>
            </span>
          </InputElm>
        </form>
      </section>
    </div>
  );
}

function SideSettings({ id }: { id: string }) {
  const rulesList = [
    { name: "Default", content: "default" },
    { name: "Custom", content: "custom" },
  ];

  const [selectedRule, setSelectedRule] = useState(rulesList[0].name);

  const { data, isLoading, error } = useBlockRules(id);

  const { mutate } = useMutation(async (rules: typeof data) => {
    await axios.put(`/cmd/block/${id}`, { rules });
  });

  const handleNumber = (e: FormEvent<HTMLInputElement>) => {
    if (data) {
      mutate({ ...data, [e.currentTarget.id]: e.currentTarget.value });
    }
  };

  const handleSwitch = (id: string, value: boolean) => {
    if (data) {
      mutate({ ...data, [id]: value });
    }
  };

  const handleRuleNameChange = (value: string) => {
    setSelectedRule(value);
    console.log(value);
  };

  if (isLoading) return <Loading />;

  if (error || !data) return <Error />;

  return (
    <div className="flex flex-col w-full overflow-hidden h-fit">
      <Island className="relative flex flex-col justify-center px-5 h-1/3">
        <span className="w-5 h-5 absolute top-1 right-1">
          <GearIcon />
        </span>
        <div className="">
          <label className="font-bold inline mx-5" htmlFor="nested">
            Allow to Be Nested
          </label>
          <SwitchToggle
            id="nested"
            value={data.nested}
            onChange={handleSwitch}
          />
        </div>
        <div>
          <label className="font-bold inline mx-5" htmlFor="hasLikes">
            Allow Likes
          </label>
          <SwitchToggle
            id="hasLikes"
            value={data.hasLikes}
            onChange={handleSwitch}
          />
        </div>
        <div>
          <label className="font-bold inline mx-5" htmlFor="hasComments">
            Allow Comments
          </label>
          <SwitchToggle
            id="hasComments"
            value={data.hasComments}
            onChange={handleSwitch}
          />
        </div>
        <div>
          <label className="font-bold inline mx-5" htmlFor="commentsHasLikes">
            Allow Comments To have Likes
          </label>
          <SwitchToggle
            id="commentsHasLikes"
            value={data.commentsHasLikes}
            onChange={handleSwitch}
          />
        </div>
        <div>
          <label className="font-bold inline mx-5" htmlFor="commentsEditable">
            Allow Comments To Be Editable
          </label>
          <SwitchToggle
            id="commentsEditable"
            value={data.commentsEditable}
            onChange={handleSwitch}
          />
        </div>
        <div>
          <label className="font-bold inline mx-5" htmlFor="commentsMaxNested">
            Comments MaxNested (0 for no nesting)
          </label>
          <input
            type="number"
            defaultValue={data.commentsMaxNested}
            onChange={handleNumber}
            id="commentsMaxNested"
            className="p-1 rounded shadow border border-gray-100 text-black w-7 h-5"
          />
        </div>
        <div className="w-8/12 h-px bg-gray-500 my-5 mx-auto"></div>
        <div className="flex flex-row justify-evenly my-1">
          <h3>Auto-Fill Rules: </h3>
          <Select
            selected={selectedRule}
            items={rulesList}
            className="shadow-md py-1 px-3 rounded-full hover:bg-s-gopher hover:text-white"
            onChange={handleRuleNameChange}
            ariaLabel="Select A Rule Set"
          />
        </div>
      </Island>
      <AuthorInfo />
    </div>
  );
}

function AuthorInfo() {
  const user = {
    id: "bla-563-468",
    name: "Jihed Mastouri",
    image: "https://via.placeholder.com/150",
    role: "Admin",
  };

  return (
    <>
      <div className="flex mt-2">
        <span className="w-5 h-5">
          <PersonIcon />
        </span>
        <h3 className="font-bold">Author:</h3>
      </div>
      <Link
        className="flex justify-start items-center my-1 p-1 ml-5 max-w-[300px] shadow-md border-gray-500 rounded-full hover:bg-s-gopher-l hover:text-gray-700"
        to={`/users/${user.id}`}
      >
        <img className="w-10 h-10 rounded-full ml-2 mr-5" src={user.image} />
        <div className="mr-5 overflow-hidden overflow-ellipsis whitespace-nowrap">
          <p className="text-s-grass font-bold">{user.name}</p>
          <p className="text-sm font-thin">{user.role}</p>
        </div>
      </Link>
    </>
  );
}

function MainImage() {
  return (
    <Island className="p-1">
      <img
        className="object-cover w-full max-h-[300px]"
        src="https://via.placeholder.com/150"
      />
      <div className="flex items-center justify-end w-11/12 h-10 m-auto">
        <button className="w-7 h-7 border rounded-lg p-1 hover:bg-s-gopher hover:text-white">
          <Pencil1Icon />
        </button>
        <button className="ml-2 w-7 h-7 border rounded-lg p-1 hover:bg-s-gopher hover:text-white">
          <UploadIcon />
        </button>
      </div>
    </Island>
  );
}
