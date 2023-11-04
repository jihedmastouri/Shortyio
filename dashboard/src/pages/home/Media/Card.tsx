import * as AspectRatio from "@radix-ui/react-aspect-ratio";
import Zoom from "./Zoom";
import { TrashIcon, ZoomInIcon } from "@radix-ui/react-icons";
import { Row } from "@tanstack/react-table";
import ConfirmAction from "@ui/ConfirmAction";
import { motion } from "framer-motion";
import { useState } from "react";
import { Media } from "./media";

type Props = {
  data: Row<Media>;
  index: number;
};

const CardImage = ({ data, index }: Props) => {
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
      className="overflow-hidden rounded min-w-[350px] w-[30%] m-1 shadow-black"
    >
      <AspectRatio.Root ratio={16 / 9}>
        <motion.div
          initial={{ opacity: 0 }}
          animate={isHovered ? { opacity: 1 } : { opacity: 0 }}
          transition={{ type: "spring", stiffness: 100 }}
          className="flex w-full h-full items-end absolute text-center"
        >
          <p className="bg-white max-w-full mx-1 mb-1 px-2 rounded shadow opacity-75 hover:opacity-100 whitespace-nowrap overflow-hidden overflow-ellipsis">
            {data.original.title}
          </p>
        </motion.div>
        <motion.div
          initial={{ opacity: 0 }}
          animate={isHovered ? { opacity: 1 } : { opacity: 0 }}
          transition={{ type: "spring", stiffness: 100 }}
          className="flex w-full justify-end absolute"
        >
          <div className="flex flex-col justify-center bg-white mr-2 my-2 shadow rounded opacity-75 hover:opacity-100 ">
            <Zoom
              img={data.original.image}
              alt={data.original.alt}
              className="w-5 h-5 hover:bg-gray-200 rounded bg-white"
            >
              <ZoomInIcon />
            </Zoom>
            <ConfirmAction
              message="You are about to delete this image"
              fun={() => {}}
              className="w-5 h-5 hover:bg-red-200 rounded"
            >
              <TrashIcon />
            </ConfirmAction>
          </div>
        </motion.div>
        <img className="object-cover w-full h-full" src={data.original.image} />
      </AspectRatio.Root>
    </motion.div>
  );
};

export default CardImage;
