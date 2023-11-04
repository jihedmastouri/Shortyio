import * as Label from "@radix-ui/react-label";
import * as Separator from "@radix-ui/react-separator";
import { ReactNode } from "react";
import Island from "../../components/Island";

const InputElm = ({
  name,
  id,
  children,
}: {
  name: string;
  id: string;
  children: ReactNode;
}) => {
  return (
    <Island className="flex items-center">
      <Label.Root className="font-bold text-lg inline" htmlFor={id}>
        {name}
      </Label.Root>
      <Separator.Root
        className="bg-gray-800 rounded-sm opacity-25 w-[3px] h-7 ml-auto mr-16"
        decorative
        orientation="vertical"
      />
      {children}
    </Island>
  );
};

export default InputElm;
