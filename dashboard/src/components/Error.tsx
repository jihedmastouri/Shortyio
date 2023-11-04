import { MinusCircledIcon } from "@radix-ui/react-icons";

export default function Error() {
  return (
    <div
      className="w-full h-screen z-30 overflow-hidden bg-red-700 opacity-75 flex flex-col items-center justify-center"
    >
      <div className="h-12 w-12 mb-2 text-white">
        <MinusCircledIcon />
      </div>
      <h2 className="text-center text-white text-xl font-semibold">Error</h2>
      <p className="w-1/3 text-center text-white">
        Something went wrong :(
      </p>
    </div>
  )
}
