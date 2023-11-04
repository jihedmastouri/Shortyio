import * as AspectRatio from "@radix-ui/react-aspect-ratio";
import { ZoomInIcon } from "@radix-ui/react-icons";
import { ColumnDef, createColumnHelper } from "@tanstack/react-table";
import AddEntries from "@ui/AddEntries";
import Header from "../components/Header";
import Card from "./Card";
import data from "./imges.json";
import { Media } from "./media";
import Zoom from "./Zoom";

const columnHelper = createColumnHelper<Media>();

const columns: ColumnDef<Media, any>[] = [
  columnHelper.display({
    header: "The image",
    id: "img",
    cell: (info) => (
      <Zoom
        img={info.row.original.image}
        alt={info.row.original.alt}
        className="h-20 w-20 overflow-hidden"
      >
        <div className="max-h-full w-20 mx-auto">
          <AspectRatio.Root ratio={1 / 1} className="rounded">
            <img
              className="w-full h-full object-cover"
              src={info.row.original.image}
              alt={info.row.original.alt}
            />
          </AspectRatio.Root>
        </div>
      </Zoom>
    ),
    footer: "an image",
  }),
  columnHelper.accessor("title", {
    header: "Title",
    cell: (info) => <p className="w-10">{info.getValue()}</p>,
    footer: "Title",
    enableSorting: true,
    size: 300,
  }),
  columnHelper.accessor("alt", {
    header: "Description",
    cell: (info) => <p className="w-50">{info.getValue()}</p>,
    footer: "Title",
    enableSorting: true,
    size: 300,
  }),
  columnHelper.accessor("dateCreated", {
    header: "Date",
    cell: (info) => new Date(info.getValue()).toLocaleString(),
    footer: "Date",
    enableSorting: true,
    size: 300,
  }),
  columnHelper.display({
    id: "action",
    header: "",
    size: 10,
    cell: () => (
      <span className="w-5 h-5">
        <ZoomInIcon />
      </span>
    ),
  }),
];

const index = () => {
  return (
    <div className="w-full h-full overflow-x-hidden">
      <div className="mt-2 mb-5 mx-5 flex justify-between">
        <Header />
        <AddEntries entriesNumber={20} entriesType="Images" subFun={() => {}} />
      </div>
      {
        // <Table data={data} columns={columns} layoutType="grid" GridCard={Card} />
      }
    </div>
  );
};

export default index;
