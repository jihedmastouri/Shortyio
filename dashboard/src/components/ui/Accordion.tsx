import { PlusIcon } from "@radix-ui/react-icons";

type Props = {
  entriesNumber: number;
  entriesType: string;
  subFun: Function;
};

const Accordion = ({ entriesNumber, subFun, entriesType }: Props) => {
  return (
    <div className="flex items-center border border-gray-500 shadow-lg rounded h-8 w-fit self-end mb-2">
      <h3 className="text-lg font-bold px-10">
        {entriesNumber} {entriesType}
      </h3>
      <button
        className="bg-black opacity-80 h-8 w-8 text-white rounded-r"
        onClick={() => subFun()}
      >
        <PlusIcon />
      </button>
    </div>
  );
};

export default Accordion;
