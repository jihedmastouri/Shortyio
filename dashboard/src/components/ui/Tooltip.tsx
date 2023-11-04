import { ReactNode } from "react";
import * as Tooltip from "@radix-ui/react-tooltip";

type Props = {
  duration?: number;
  children: ReactNode;
  content: string;
  direction?: "top" | "right" | "bottom" | "left";
};

export default function appTooltip({
  duration = 200,
  children,
  content,
  direction = "top",
}: Props) {
  return (
    <Tooltip.Provider delayDuration={duration}>
      <Tooltip.Root>
        <Tooltip.Trigger asChild>{children}</Tooltip.Trigger>
        <Tooltip.Portal>
          <Tooltip.Content
            className={
              "data-[state=delayed-open]:data-[side=top]:animate-slideDownAndFade data-[state=delayed-open]:data-[side=right]:animate-slideLeftAndFade data-[state=delayed-open]:data-[side=left]:animate-slideRightAndFade data-[state=delayed-open]:data-[side=bottom]:animate-slideUpAndFade" +
              " " +
              "py-1 px-3 text-lg font-bold bg-s-gopher text-white rounded-md shadow-md will-change-[transform,opacity]"
            }
            side={direction}
          >
            {content}
            <Tooltip.Arrow className="fill-s-gopher" />
          </Tooltip.Content>
        </Tooltip.Portal>
      </Tooltip.Root>
    </Tooltip.Provider>
  );
}
