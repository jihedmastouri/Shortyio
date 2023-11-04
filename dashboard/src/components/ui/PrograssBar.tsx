import * as Progress from "@radix-ui/react-progress";

type Props = {
  progress: number;
};

export default function PrograssBar({ progress }: Props) {
  return (
    <Progress.Root
      className="bg-gray-400 relative overflow-hidden rounded translate-z-0 w-full h-full"
      value={progress}
    >
      <Progress.Indicator
        className="bg-blue-200 w-full h-full"
        style={{
          transform: `translateX(-${100 - progress}%)`,
          transition: "transform 660ms cubic-bezier(0.65, 0, 0.35, 1)",
        }}
      />
    </Progress.Root>
  );
}
