import { useContext, HTMLProps, useRef } from "react";
import * as Toolbar from "@radix-ui/react-toolbar";
import {
  ChevronLeftIcon,
  ChevronRightIcon,
  DoubleArrowLeftIcon,
  DoubleArrowRightIcon,
} from "@radix-ui/react-icons";

import Select from "@ui/Select";
import { Table as TableContext } from "./Contexts";

// TODO: SWITCH TO MANUAL PAGINATION
export default function Paginator(props: HTMLProps<HTMLDivElement>) {
  const table = useContext(TableContext)?.table!;
  const tableSize = useRef(10);

  const hoverStyle = "hover:text-white hover:bg-s-gopher hover:border-s-grass";
  const arrowStyle =
    "border border-gray-500 rounded-full w-4 h-4 bg-gray-50 flex items-center justify-center";

  return (
    <div className="select-none sticky top-2 z-10 bg-white mt-1 p-1 w-fit mx-auto rounded-full shadow-md text-sm">
      <Toolbar.Root
        className={
          props.className
            ? props.className
            : "flex justify-center items-center gap-2 w-fit mx-auto px-2 py-1 rounded-full"
        }
        aria-label="Pagination"
      >
        <span className="flex items-center gap-1">
          Go to page:
          <input
            type="number"
            defaultValue={table.getState().pagination.pageIndex + 1}
            onChange={(e) => {
              const page = e.target.value ? Number(e.target.value) - 1 : 0;
              table.setPageIndex(page);
            }}
            className="border p-1 rounded w-10 h-5"
          />
        </span>

        <Toolbar.Separator className="w-1 bg-gray-500 mx-3" />

        <div className="flex justify-center items-center gap-2">
          <Toolbar.Button
            className={arrowStyle + " " + hoverStyle}
            onClick={() => table.setPageIndex(0)}
            disabled={!table.getCanPreviousPage()}
          >
            <DoubleArrowLeftIcon />
          </Toolbar.Button>
          <Toolbar.Button
            className={arrowStyle + " " + hoverStyle}
            onClick={() => table.previousPage()}
            disabled={!table.getCanPreviousPage()}
          >
            <ChevronLeftIcon />
          </Toolbar.Button>
          <span className="flex items-center gap-1">
            <div>Page</div>
            <strong>
              {table.getState().pagination.pageIndex + 1} of{" "}
              {table.getPageCount()}
            </strong>
          </span>
          <Toolbar.Button
            className={arrowStyle + " " + hoverStyle}
            onClick={() => table.nextPage()}
            disabled={!table.getCanNextPage()}
          >
            <ChevronRightIcon />
          </Toolbar.Button>
          <Toolbar.Button
            className={arrowStyle + " " + hoverStyle}
            onClick={() => table.setPageIndex(table.getPageCount() - 1)}
            disabled={!table.getCanNextPage()}
          >
            <DoubleArrowRightIcon />
          </Toolbar.Button>
        </div>

        <Toolbar.Separator className="w-6 bg-gray-500 mx-3" />

        <Select
          ariaLabel="Show 10"
          className={
            "rounded-md p-1 bg-gray-50 shadow shadow-gray-200 " + hoverStyle
          }
          selected={tableSize.current}
          onChange={(value) => {
            table.setPageSize(Number(value));
            tableSize.current = Number(value);
          }}
          items={[10, 20, 30, 40, 50].map((pageSize) => ({
            name: `${pageSize}`,
            content: `Show ${pageSize}`,
          }))}
        />
      </Toolbar.Root>
    </div>
  );
}
