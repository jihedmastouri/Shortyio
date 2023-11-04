import { ColumnDef, createColumnHelper } from "@tanstack/react-table";
import Table from "@ui/Table";
import { Link } from "react-router-dom";
import data from "./comment.json";

type CommentType = {
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
  parent?: number;
  last_updated: string;
};

const columnHelper = createColumnHelper<CommentType>();

const columns: ColumnDef<CommentType, any>[] = [
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

const Comments = () => {
  return <Table data={data} columns={columns} />;
};

export default Comments;
