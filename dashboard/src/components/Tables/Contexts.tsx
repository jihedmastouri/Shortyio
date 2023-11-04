import { createContext } from "react";
import {
  ColumnResizeMode,
  Table as tanstackTable,
} from "@tanstack/react-table";

type TableType<T extends object> = {
  table: tanstackTable<T>;
};

const Table = createContext<TableType<any> | null>(null);
const ColumnRM = createContext<ColumnResizeMode>("onChange");

export { Table, ColumnRM };
