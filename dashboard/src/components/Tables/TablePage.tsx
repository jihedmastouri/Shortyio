import { ReactNode, useContext } from "react";

import DebouncedInput from "@ui/Search";
import Paginator from "./Paginator";
import { Table as TableContext } from "./Contexts";

type Props = {
  children: ReactNode;
};

export default function TablePage({ children }: Props) {
  const table = useContext(TableContext)?.table!;

  return (
    <div className="flex flex-col justify-between w-full h-fit">
      <DebouncedInput
        value={table.getState().globalFilter || ""}
        onChange={(value) => table.setGlobalFilter(String(value))}
        className="p-2 font-lg shadow border border-gray-300 rounded-full w-1/3 ml-5"
        placeholder="Search all columns..."
      />
      <Paginator />
      {children}
    </div>
  );
}
