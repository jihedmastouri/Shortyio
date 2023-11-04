import { useContext } from "react";

import { flexRender } from "@tanstack/react-table";
import { motion } from "framer-motion";

import { Table as TableContext, ColumnRM as ColumnRMContext } from "./Contexts";
import { ArrowDownIcon, ArrowUpIcon } from "@radix-ui/react-icons";

export default function Base(props: { tr?: string; td?: string }) {
  const table = useContext(TableContext)?.table!;
  const columnResizeMode = useContext(ColumnRMContext);

  return (
    <>
      <table className=" mx-auto overflow-x-scroll border-spacing-1 border-separate">
        <thead className="h-10 text-s-grass text-lg">
          {table.getHeaderGroups().map((headerGroup) => (
            <tr key={headerGroup.id}>
              <th className="text-sm"></th>
              {headerGroup.headers.map((header) => (
                <th
                  key={header.id}
                  className="pl-2 cursor-pointer select-none"
                  onClick={header.column.getToggleSortingHandler()}
                >
                  <span className="flex items-center justify-center">
                    {header.isPlaceholder
                      ? null
                      : flexRender(
                          header.column.columnDef.header,
                          header.getContext()
                        )}
                    <span className="h-4 w-4 m-4">
                      {{
                        asc: <ArrowUpIcon />,
                        desc: <ArrowDownIcon />,
                      }[header.column.getIsSorted() as string] ?? null}
                    </span>
                    <div
                      {...{
                        onMouseDown: header.getResizeHandler(),
                        onTouchStart: header.getResizeHandler(),
                        className: `resizer ${
                          header.column.getIsResizing() ? "isResizing" : ""
                        }`,
                        style: {
                          transform: header.column.getIsResizing()
                            ? `translateX(${
                                columnResizeMode == "onEnd" &&
                                table.getState().columnSizingInfo.deltaOffset
                              }px)`
                            : "",
                        },
                      }}
                    />
                  </span>
                </th>
              ))}
            </tr>
          ))}
        </thead>

        <tbody>
          {table.getRowModel().rows.map((row) => (
            <tr key={row.id} className={props.tr}>
              <motion.td
                initial={{ x: 10, opacity: 0 }}
                animate={{ x: -10, opacity: 1 }}
                transition={{
                  type: "spring",
                  stiffness: 100,
                  delay: 0.1,
                }}
                className="text-s-grass font-bold text-sm mx-auto text-center"
              >
                {" "}
                {row.index + 1}{" "}
              </motion.td>
              {row.getVisibleCells().map((cell, index) => (
                <motion.td
                  className={props.td}
                  key={cell.id}
                  initial={{ x: 10, opacity: 0 }}
                  animate={{ x: -10, opacity: 1 }}
                  transition={{
                    type: "spring",
                    stiffness: 100,
                    delay: 0.1 * index,
                  }}
                >
                  {flexRender(cell.column.columnDef.cell, cell.getContext())}
                </motion.td>
              ))}
            </tr>
          ))}
        </tbody>

        <tfoot className="border-t bg-gray-50 font-light text-sm text-slate-400">
          {table.getFooterGroups().map((footerGroup) => (
            <tr key={footerGroup.id}>
              <th className="text-sm"></th>
              {footerGroup.headers.map((header) => (
                <th key={header.id} colSpan={header.colSpan}>
                  {header.isPlaceholder
                    ? null
                    : flexRender(
                        header.column.columnDef.footer,
                        header.getContext()
                      )}
                </th>
              ))}
            </tr>
          ))}
        </tfoot>
      </table>
    </>
  );
}
