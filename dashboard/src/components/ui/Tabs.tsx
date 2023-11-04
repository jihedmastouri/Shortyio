import * as Tabs from "@radix-ui/react-tabs";
import { AnimatePresence, motion } from "framer-motion";
import { ReactNode } from "react";

type Props = {
  tabNames: string[];
  tabs: ReactNode[];
};

const MainTabs = ({ tabs, tabNames }: Props) => {
  return (
    <Tabs.Root className="flex flex-col w-full" defaultValue="tab1">
      <Tabs.List
        className="flex justify-end mb-2"
        aria-label="Manage your account"
      >
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
      <hr className="mb-5 h-1 opacity-25 bg-gray-500 " />
      {tabs.map((el, index) => (
        <Tabs.Content
          className="flex flex-col h-full"
          key={"ptab-content-" + index}
          value={"tab" + (index + 1).toString()}
        >
          <AnimatePresence>
            <motion.div
              // key={selectedTab ? selectedTab.label : "empty"}
              initial={{ x: 10, opacity: 0 }}
              animate={{ x: 0, opacity: 1 }}
              exit={{ y: -10, opacity: 0 }}
              transition={{ type: "spring", stiffness: 50 }}
            >
              {el}
            </motion.div>
          </AnimatePresence>
        </Tabs.Content>
      ))}
    </Tabs.Root>
  );
};

export default MainTabs;
