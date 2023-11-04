import * as DropdownMenu from "@radix-ui/react-dropdown-menu";
import { ChevronDownIcon, ChevronUpIcon } from "@radix-ui/react-icons";
import { ReactNode, useState } from "react";

type Props = {
  children: ReactNode;
  content: ReactNode[];
  role?: string;
};

function DropDown({ content, role, children }: Props) {
  const [isOpen, setIsOpen] = useState(false);
  const toggleState = () => {
    setIsOpen((old) => !old);
  };
  return (
    <DropdownMenu.Root onOpenChange={toggleState}>
      <DropdownMenu.Trigger asChild>
        <button
          className="rounded-lg border-black shadow h-[90%] my-auto flex items-center justify-between px-3 py-1"
          aria-label={role}
        >
          {children}
          {isOpen ? (
            <ChevronUpIcon className="w-5" />
          ) : (
            <ChevronDownIcon className="w-5" />
          )}
        </button>
      </DropdownMenu.Trigger>
      <DropdownMenu.Content
        className="bg-white p-3 rounded-sm shadow-white border-zinc-50 border z-40"
        sideOffset={5}
      >
        {content.map((el, index) => (
          <DropdownMenu.Item key={"drop-down-" + index}>{el}</DropdownMenu.Item>
        ))}
      </DropdownMenu.Content>
    </DropdownMenu.Root>
  );
}

export default DropDown;
