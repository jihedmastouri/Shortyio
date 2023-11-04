import {
  Column,
  ColumnDef,
  createColumnHelper,
  FilterFn,
  Row,
} from "@tanstack/react-table";
import Table from "@ui/Table";
import { Link } from "react-router-dom";
import data from "./like.json";

type LikeType = {
  id: number;
  user: {
    image: string;
    name: string;
    id: number;
  };
  post: {
    id: number;
    title: string;
  };
  last_updated: string;
};

function userFilter(rows: Row<LikeType>[], _: number, filterValue: string) {
  return filterValue.length === 0
    ? rows
    : rows.filter((row) =>
        filterValue.includes(String(row.original.user.name))
      );
}

const columnHelper = createColumnHelper<LikeType>();

const columns: ColumnDef<LikeType, any>[] = [
  columnHelper.accessor("user", {
    header: "User",
    cell: (info) => (
      <Link
        className="flex items-center text-xl"
        to={"/user/" + info.getValue().id}
      >
        <img
          className="w-5 h-5 mr-2 rounded-full"
          src={info.getValue().image}
          alt="Commenter Profile"
        />
        {info.getValue().name}
      </Link>
    ),
    // filter: userFilter,
	enableGlobalFilter: true,
    footer: "User",
    enableSorting: true,
    size: 300,
  }),
  columnHelper.accessor("post", {
    header: "Post",
    cell: (info) => (
      <Link
        className="flex items-center text-xl"
        to={"/user/" + info.getValue().id}
      >
        {info.getValue().title}
      </Link>
    ),
    footer: "User",
    enableSorting: true,
    size: 300,
  }),
  columnHelper.accessor("last_updated", {
    header: "Last Updated",
    cell: (info) => new Date(info.getValue()).toLocaleString(),
    footer: "Last Updated",
    enableSorting: true,
    size: 300,
  }),
];

const Likes = () => {
  return <Table data={data} columns={columns} />;
};

export default Likes;
