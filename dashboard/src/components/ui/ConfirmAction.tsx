import * as Dialog from "@radix-ui/react-dialog";
import { Cross2Icon } from "@radix-ui/react-icons";
import { HTMLAttributes, ReactNode } from "react";

type Props = {
  message: string;
  fun: Function;
  children: ReactNode;
};

const ConfirmAction = ({
  message,
  fun,
  children,
  className,
}: Props & HTMLAttributes<HTMLDivElement>) => {
  return (
    <Dialog.Root>
      <Dialog.Trigger className={className}> {children}</Dialog.Trigger>
      <Dialog.Portal>
        <Dialog.Overlay className="z-50 bg-gray-600 fixed inset-0 opacity-30" />
        <Dialog.Content
          className={
            "w-1/3 h-fit my-5 px-6 bg-white rounded-lg shadow " +
            " z-50 fixed top-[50%] left-[50%] translate-x-[-50%] translate-y-[-50%]"
          }
        >
          <Dialog.Close asChild>
            <button
              className="bg-red-300 text-white w-5 h-5 absolute right-0 top-0 m-2 rounded-full"
              aria-label="Close"
            >
              <Cross2Icon />
            </button>
          </Dialog.Close>

          <Dialog.Title className="font-bold text-xl my-5">
            Please Confirm This Action
          </Dialog.Title>
          <Dialog.Description className="text-gray-500 my-5">
            {message}
          </Dialog.Description>

          <div className="flex justify-end mb-5">
            <Dialog.Close asChild>
              <button
                className="bg-s-gopher-l hover:bg-s-gopher hover:text-white px-3 py-2 rounded"
                onClick={() => fun()}
              >
                Yes, I am Sure
              </button>
            </Dialog.Close>
          </div>
        </Dialog.Content>
      </Dialog.Portal>
    </Dialog.Root>
  );
};

export default ConfirmAction;
