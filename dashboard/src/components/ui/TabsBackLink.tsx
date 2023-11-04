import React, { ReactNode } from "react";
import * as Tabs from "@radix-ui/react-tabs";
import { Link } from "react-router-dom";
import { ArrowLeftIcon } from "@radix-ui/react-icons";

type Props = {
  tabNames: string[];
  tabs: ReactNode[];
  backLink: {
    name: string;
    link: string;
  };
};

const MainTabs = ({ tabs, tabNames, backLink }: Props) => {
  return (
    <Tabs.Root className="flex flex-col w-full" defaultValue="tab1">
      <Tabs.List
        className="flex justify-end mb-5"
        aria-label="Manage your account"
      >
        <Link
          to={backLink.link}
          className="flex items-center text-xl px-2 group hover:underline mr-auto"
        >
          <span className="h-8 w-10 px-2 flex group-hover:-translate-x-2">
            <ArrowLeftIcon className="m-auto" />
          </span>
          {backLink.name}
        </Link>
        {tabNames.map((el, index) => (
          <Tabs.Trigger
            className="tabs"
            key={"ptab-name-" + index}
            value={"tab" + (index + 1).toString()}
          >
            {el}
          </Tabs.Trigger>
        ))}
      </Tabs.List>
      <hr className="mb-2 h-1 opacity-25 bg-gray-500 " />
      {tabs.map((el, index) => (
        <Tabs.Content
          className="flex flex-col h-full"
          key={"ptab-content-" + index}
          value={"tab" + (index + 1).toString()}
        >
          {el}
        </Tabs.Content>
      ))}
    </Tabs.Root>
  );
};

export default MainTabs;
