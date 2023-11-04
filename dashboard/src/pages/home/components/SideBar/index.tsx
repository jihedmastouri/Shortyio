import * as NavigationMenu from "@radix-ui/react-navigation-menu";
import Avatar from "@ui/Avatar";

import { NavLink } from "react-router-dom";
import { useState } from "react";

import { linkList, LinkType } from "./router";
import DropDown from "@ui/DropDown";
import AppToolTips from "@ui/Tooltip";

const LinkElement = ({ icon, name, link }: LinkType) => {
  const [isVisible, setVisible] = useState(true);

  return isVisible ? (
    <NavigationMenu.Item>
      <AppToolTips content={name} direction="right">
        <NavigationMenu.Link asChild>
          <NavLink to={link}>
            {({ isActive }) => (
              <span
                className={
                  (isActive ? "text-amber-400" : "hover:text-s-gopher") +
                  " " +
                  "w-[80%] my-5 h-10 mx-auto flex justify-center"
                }
              >
                {icon}
              </span>
            )}
          </NavLink>
        </NavigationMenu.Link>
      </AppToolTips>
    </NavigationMenu.Item>
  ) : null;
};

const SideBar = () => {
  const user: info = {
    image:
      "https://images.unsplash.com/photo-1492633423870-43d1cd2775eb?&w=128&h=128&dpr=2&q=80",
    // image: "",
    name: "salah",
    lastName: "salah",
  };

  return (
    <NavigationMenu.Root
      orientation="vertical"
      className="w-16 flex flex-col bg-white dark:bg-gray-800 shadow rounded-l-md h-full"
    >
      <NavigationMenu.List className="text-center mt-8 w-full h-full relative">
        {linkList.map((el, index) => {
          return <LinkElement {...el} key={"nav-" + index} />;
        })}
        <NavigationMenu.Item className="mt-auto">
          <NavigationMenu.Trigger className="w-5/6 mx-auto">
            <Avatar />
          </NavigationMenu.Trigger>
          <NavigationMenu.Content className="absolute top-0 left-0 w-full">
            Jeff
          </NavigationMenu.Content>
        </NavigationMenu.Item>
      </NavigationMenu.List>
    </NavigationMenu.Root>
  );
};

export default SideBar;
