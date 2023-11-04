import { ReactNode, useMemo, useState } from "react";

import {
  getCoreRowModel,
  getSortedRowModel,
  getFilteredRowModel,
  getPaginationRowModel,
  useReactTable,
  ColumnDef,
} from "@tanstack/react-table";

import { AnimatePresence } from "framer-motion";

import Base from "./Table";
import Grid from "./TableGrid";
import Page from "./TablePage";
import * as contexts from "./Contexts";

interface Props<T extends object> {
  children: ReactNode;
  data: T[];
  columns: ColumnDef<T, unknown>[];
}

function Root<T extends object>({
  children,
  data: dataProps,
  columns: columnsProps,
}: Props<T>) {
  const data = useMemo(() => dataProps, [dataProps]);
  const columns = useMemo(() => columnsProps, [columnsProps]);

  const instance = useReactTable({
    data,
    columns,
    getCoreRowModel: getCoreRowModel(),
    getFilteredRowModel: getFilteredRowModel(),
    getSortedRowModel: getSortedRowModel(),
    getPaginationRowModel: getPaginationRowModel(),
  });

  return (
    <div className="w-11/12 mx-auto">
      <div className="relative bg-none flex flex-col">
        <AnimatePresence>
          <contexts.Table.Provider value={{ table: instance }}>
              {children}
          </contexts.Table.Provider>
        </AnimatePresence>
      </div>
    </div>
  );
}

/*
// @ts-ignore
function fuzzySort(rowA, rowB, columnId) {
  let dir = 0;

  // Only sort by rank if the column has ranking information
  if (rowA.columnFiltersMeta[columnId]) {
    dir = compareItems(
      rowA.columnFiltersMeta[columnId]?.itemRank!,
      rowB.columnFiltersMeta[columnId]?.itemRank!
    );
  }

  // Provide an alphanumeric fallback for when the item ranks are equal
  return dir === 0 ? sortingFns.alphanumeric(rowA, rowB, columnId) : dir;
}

// @ts-ignore
function fuzzyFilter(row, columnId, value, addMeta) {
  // Rank the item
  const itemRank = rankItem(row.getValue(columnId), value);

  // Store the itemRank info
  addMeta({
    itemRank,
  });

  // Return if the item should be filtered in/out
  return itemRank.passed;
}
*/

export { Grid, Base, Page, Root };
