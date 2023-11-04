import * as AspectRatio from "@radix-ui/react-aspect-ratio";
import { FileIcon } from "@radix-ui/react-icons";
import { Row } from "@tanstack/react-table";
import { motion } from "framer-motion";
import { useState } from "react";
import { Link } from "react-router-dom";
import { Post } from "./post";

type Props = {
  data: Row<Post>;
  index: number;
};

export default function CardImage({ data, index }: Props) {
  // const haha = data.map(el => ({ name: el.id, value: el.getValue() }))
  const [isHovered, setisHovered] = useState(false);
  const onMouseEnter = () => {
    setisHovered(true);
  };
  const onMouseLeave = () => {
    setisHovered(false);
  };

  return (
    <motion.div
      onMouseEnter={onMouseEnter}
      onMouseLeave={onMouseLeave}
      initial={{ y: 10, opacity: 0 }}
      animate={{ y: -10, opacity: 1 }}
      transition={{ type: "spring", stiffness: 100, delay: 0.1 * index }}
      className="overflow-hidden rounded-md w-[30%] min-w-[350px] m-1 shadow-lg bg-white dark:bg-neutral-700"
    >
      <Link
        to={`${data.original.id}`}
        data-te-ripple-init
        data-te-ripple-color="light"
      >
        <AspectRatio.Root ratio={16 / 9}>
          {/*
          <img className="w-full h-full" src={data.original.image} alt="" />
          */}
          <span className="w-full h-full">
            <FileIcon />
          </span>
        </AspectRatio.Root>
        <div className="p-6">
          {/*
        <p className="mb-2 text-sm font-medium leading-tight text-neutral-500 dark:text-neutral-400">
          {data.original.last_updated}
        </p>
        {data.original.status}
      */}
          <h5 className="mb-2 text-xl font-medium leading-tight text-neutral-800 dark:text-neutral-50">
            {data.original.name}
          </h5>
          <p className="mb-4 text-base text-neutral-600 dark:text-neutral-200">
            {limitText(data.original.description, 200)}
          </p>
        </div>
      </Link>
    </motion.div>
  );
}

function limitText(text: string, limit: number) {
  if (!text) return "";
  if (text.length > limit) {
    return text.slice(0, limit) + "...";
  }
  return text;
}
