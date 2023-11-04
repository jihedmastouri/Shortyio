import { Pencil2Icon, TrashIcon } from "@radix-ui/react-icons";
import { ColumnDef } from "@tanstack/react-table";
import * as Table from "components/Tables";
import data from "./editors.json";
import ConfirmAction from "@ui/ConfirmAction";
import AddEntries from "@ui/AddEntries";

type EditorType = {
  id: number;
  image: string;
  name: string;
  data_joined: string;
  role: string;
};

const columns: ColumnDef<EditorType, any>[] = [
  {
    accessorKey: "name",
    cell: (info) => info.getValue(),
    enableSorting: true,
    size: 100,
  },
  {
    accessorKey: "image",
    enableSorting: false,
    cell: (info) => (
      <img src={info.getValue()} className="mx-auto object-cover w-20 h-20" />
    ),
    size: 100,
  },
  {
    accessorKey: "data_joined",
    header: "Date Joined",
    cell: (info) => new Date(info.getValue()).toLocaleString(),
    enableSorting: true,
    size: 100,
  },
  {
    accessorKey: "role",
    cell: (info) => info.getValue(),
    enableSorting: true,
    size: 100,
  },
  {
    accessorKey: "id",
    header: "",
    cell: (info) => <EditEditor id={info.getValue()} />,
    enableSorting: false,
    size: 50,
  },
];

const Editors = () => {
  return (
    <>
      <div className="mx-5">
        <AddEntries
          entriesNumber={3}
          entriesType={"Editors"}
          subFun={() => {}}
        />
      </div>
      <Table.Root data={data} columns={columns}>
        <Table.Base />
      </Table.Root>
    </>
  );
};

function EditEditor({ id }: { id: number }) {
  const edit = (id: number) => {
    console.log(id);
  };
  const deleteUser = (id: number) => {
    console.log(id);
  };
  // add tooltip
  return (
    <div className="mx-auto w-1/2 flex justify-around">
      <button className="w-5 h-5" onClick={() => edit(id)}>
        <Pencil2Icon />
      </button>
      <ConfirmAction
        message="You are about to delete a Customer Action"
        fun={() => deleteUser(id)}
      >
        <button className="w-5 h-5 hover:text-red-500">
          <TrashIcon />
        </button>
      </ConfirmAction>
    </div>
  );
}

export default Editors;
