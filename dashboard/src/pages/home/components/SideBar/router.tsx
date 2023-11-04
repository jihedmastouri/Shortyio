import { ReactNode } from "react";
import {
  ChatBubbleIcon,
  CodeIcon,
  FileTextIcon,
  GearIcon,
  HomeIcon,
  ImageIcon,
  MixIcon,
  PersonIcon,
  RocketIcon,
} from "@radix-ui/react-icons";

export interface LinkType {
  icon: ReactNode;
  name: string;
  link: string;
  authLevel: number;
  siteType?: string;
}

const role = 0;
const siteType = "blog";

const l: LinkType[] = [
  {
    name: "Home",
    icon: <HomeIcon />,
    link: "/",
    authLevel: 0,
  },
  {
    name: "Posts",
    icon: <FileTextIcon />,
    link: "/posts",
    authLevel: 0,
    siteType: "blog",
  },
  {
    name: "Media",
    icon: <ImageIcon />,
    link: "/media",
    authLevel: 0,
    siteType: "blog",
  },
  {
    name: "Users",
    icon: <PersonIcon />,
    link: "users",
    authLevel: 0,
  },
  {
    name: "Likes & Comments",
    icon: <ChatBubbleIcon />,
    link: "interactions",
    authLevel: 0,
    siteType: "blog",
  },
  {
    name: "Settings",
    icon: <GearIcon />,
    link: "settings",
    authLevel: 0,
  },
  {
    name: "Developer Settings",
    icon: <CodeIcon />,
    link: "dev",
    authLevel: 0,
  },
  {
    name: "Plugins",
    icon: <MixIcon />,
    link: "settings",
    authLevel: 0,
  },
  {
    name: "Go Pro!",
    icon: <RocketIcon />,
    link: "settings",
    authLevel: 0,
  },
];

export const linkList = (() => {
  return l.filter((el) => {
    if (el.authLevel <= role) {
      if (el.siteType === undefined || el.siteType === siteType) return el;
    }
  });
})();
