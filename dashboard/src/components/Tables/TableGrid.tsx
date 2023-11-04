import { ElementType, ReactNode, useContext, useState } from "react";
import { Row } from "@tanstack/react-table";
import { Table as TableContext } from "./Contexts";

import * as ToggleGroup from "@radix-ui/react-toggle-group";
import { DashboardIcon, TableIcon } from "@radix-ui/react-icons";

type Props = {
  children: ReactNode;
  layoutDefault?: "GRID" | "TABLE";
  GridCard: ElementType<{
    data: Row<any>;
    index: number;
  }>;
};

export default function TableGrid({
  children,
  GridCard,
  layoutDefault,
}: Props) {
  const table = useContext(TableContext)?.table!;

  const [layoutType, setLayoutType] = useState<string>(layoutDefault || "GRID");

  const iconStyleBase =
    "w-10 items-center justify-center leading-4 hover:text-gray-500 cursor-pointer focus:text-s-grass" +
    "transition-colors duration-200 ease-in-out";

  const iconStyle = (val: string) =>
    iconStyleBase + " " + (layoutType === val ? "text-s-grass" : "text-base");

  return (
    <>
      <div className="bg-white shadow-md self-end z-10 p-1 rounded-full">
        <ToggleGroup.Root
          className="flex justify-center rounded-full w-fit h-8 p-1"
          type="single"
          defaultValue={layoutType}
          onValueChange={(val) => setLayoutType(val)}
          aria-label="Text alignment"
        >
          <ToggleGroup.Item
            className={iconStyle("GRID")}
            value="GRID"
            aria-label="Grid Layout"
          >
            <DashboardIcon />
          </ToggleGroup.Item>
          <ToggleGroup.Item
            className={iconStyle("TABLE")}
            value="TABLE"
            aria-label="Table Layout"
          >
            <TableIcon />
          </ToggleGroup.Item>
        </ToggleGroup.Root>
      </div>
      <div className="flex justify-between w-full h-fit">
        <div className="flex flex-col justify-center w-11/12 mx-auto mb-1">
          <div className="flex justify-end w-full">
            {layoutType === "TABLE" ? (
              children
            ) : (
              <div className="flex w-full gap-2 flex-wrap mt-4">
                {table.getRowModel().rows.map((row, index) => (
                  <GridCard key={row.id} data={row} index={index} />
                ))}
              </div>
            )}
          </div>
        </div>
      </div>
    </>
  );
}
