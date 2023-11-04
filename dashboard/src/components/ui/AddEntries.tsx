import { PlusIcon } from "@radix-ui/react-icons";
import * as Dialog from "@radix-ui/react-dialog";
import { ReactNode } from "react";

type Props = {
  entriesNumber: number;
  entriesType: string;
  children: ReactNode;
  open: boolean;
  onOpenChange: (open: boolean) => void;
};

const Accordion = ({ entriesNumber, children, entriesType, open, onOpenChange}: Props) => {
  return (
    <Dialog.Root open={open} onOpenChange={onOpenChange}>
      <Dialog.Trigger asChild>
        <div className="flex items-center border border-gray-500 shadow-lg rounded h-8 w-fit self-end mb-2">
          <h3 className="text-lg font-bold px-10">
            {entriesNumber} {entriesType}
          </h3>
          <button className="bg-black opacity-80 h-8 w-8 text-white rounded-r">
            <PlusIcon />
          </button>
        </div>
      </Dialog.Trigger>
      <Dialog.Portal>
        <Dialog.Overlay className="fixed inset-0 bg-black bg-opacity-50 z-20" />
        <Dialog.Content className="absolute top-1/2 left-1/2 transform -translate-x-1/2 -translate-y-1/2 bg-white rounded shadow-lg p-5 z-30">
          <Dialog.Title className="text-lg font-bold mb-5">
            Add a new {entriesType}
          </Dialog.Title>
          <Dialog.Description className="text-sm mb-5">
            Lorem ipsum dolor sit amet consectetur adipisicing elit. Quisquam
            voluptatum, quibusdam, quia, quod voluptates voluptatem quos
            voluptatibus quae doloribus quas natus. Quisquam voluptatum,
          </Dialog.Description>
          {children}
          <Dialog.Close asChild>
            <button className="bg-red-300 text-white">Close</button>
          </Dialog.Close>
        </Dialog.Content>
      </Dialog.Portal>
    </Dialog.Root>
  );
};

export default Accordion;
