import * as RadPopover from "@radix-ui/react-popover";
import { Cross2Icon } from "@radix-ui/react-icons";
import { ReactNode } from "react";

type Props = {
  chidlren: ReactNode[];
  content: ReactNode;
};

const Popover = ({ chidlren, content }: Props) => (
  <RadPopover.Root>
    <RadPopover.Trigger asChild>{chidlren}</RadPopover.Trigger>
    <RadPopover.Portal>
      <RadPopover.Content
        className="rounded p-5 w-[260px] bg-white will-change-[transform,opacity] data-[state=open]:data-[side=top]:animate-slideDownAndFade data-[state=open]:data-[side=right]:animate-slideLeftAndFade data-[state=open]:data-[side=bottom]:animate-slideUpAndFade data-[state=open]:data-[side=left]:animate-slideRightAndFade"
        sideOffset={5}
      >
        {content}
        <RadPopover.Close
          className="rounded-full h-[25px] w-[25px] inline-flex items-center justify-center text-violet11 absolute top-[5px] right-[5px] hover:bg-violet4 focus:shadow-[0_0_0_2px] focus:shadow-violet7 outline-none cursor-default"
          aria-label="Close"
        >
          <Cross2Icon />
        </RadPopover.Close>
        <RadPopover.Arrow className="fill-white" />
      </RadPopover.Content>
    </RadPopover.Portal>
  </RadPopover.Root>
);

export default Popover;
