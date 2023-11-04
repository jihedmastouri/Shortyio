import { ColumnDef, createColumnHelper } from "@tanstack/react-table";
import * as Table from "components/Tables";
import { Link } from "react-router-dom";
import { ExternalLinkIcon, TrashIcon } from "@radix-ui/react-icons";
import Header from "../components/Header";
import AddEntries from "@ui/AddEntries";
import { Post } from "./post";
import Card from "./Card";
import ConfirmAction from "@ui/ConfirmAction";
import axios from "axios";
import {
  QueryClient,
  useMutation,
  useQuery,
  useQueryClient,
} from "@tanstack/react-query";
import Loading from "components/Loading";
import Error from "components/Error";
import { BlockRulesSchema } from "./Edit/utils";
import { useMemo, useState } from "react";
import { z } from "zod";

const columnHelper = createColumnHelper<Post>();

function getColumns(queryClient: QueryClient) {
  const columns: ColumnDef<Post, any>[] = [
    columnHelper.accessor("name", {
      header: "Title",
      cell: (info) => info.getValue(),
      footer: "Title",
      enableSorting: true,
      size: 450,
      maxSize: 450,
    }),
    columnHelper.accessor("description", {
      header: "Description",
      cell: (info) => info.getValue(),
      enableSorting: true,
      footer: "Description",
      size: 100,
      maxSize: 150,
    }),
    columnHelper.display({
      id: "action",
      header: "",
      size: 100,
      maxSize: 450,
      cell: (info) => (
        <span className="flex justify-center">
          <ConfirmAction
            message="You are about to delete this image"
            fun={() => {
              axios.delete(`cmd/block/${info.row.original.id}`).then((res) => {
                console.log(res);
                queryClient.setQueryData(["posts"], (old: any) =>
                  old.filter((e: any) => e.id !== info.row.original.id)
                );
              });
            }}
            className="w-5 h-5 hover:fill-red-200"
          >
            <TrashIcon />
          </ConfirmAction>
          <Link
            className="w-5 h-5 flex items-center"
            to={`${info.row.original.id}`}
          >
            <ExternalLinkIcon />
          </Link>
        </span>
      ),
    }),
  ];
  return columns;
}

export default function () {
  const { data, error, isLoading } = useQuery(["posts"], async () => {
    return axios.get("public/block/search").then((res) => {
      //FIXME: without pagination
      return res.data.metas;
    });
  });

  const [open, setOpen] = useState(false);
  const queryClient = useQueryClient();
  const columns = useMemo(() => getColumns(queryClient), []);

  return (
    <div className="w-full">
      <div className="mt-2 mb-5 mx-5 flex justify-between">
        <Header />
        <AddEntries
          entriesNumber={isLoading ? "..." : data ? data.length : 0}
          entriesType="Posts"
          open={open}
          onOpenChange={(open) => setOpen(open)}
        >
          <AddPost setOpen={setOpen} />
        </AddEntries>
      </div>
      {isLoading ? (
        <Loading />
      ) : error ? (
        <Error />
      ) : (
        <>
          <Table.Root data={data} columns={columns}>
            <Table.Page>
              <Table.Grid GridCard={Card} layoutDefault="TABLE">
                <Table.Base td="bg-white p-2 rounded border border-gray-200 shadow whitespace-nowrap overflow-hidden overflow-ellipsis max-w-[500px]" />
              </Table.Grid>
            </Table.Page>
          </Table.Root>
        </>
      )}
    </div>
  );
}

// crate a form to add entries
function AddPost({ setOpen }: { setOpen: (b: boolean) => void }) {
  const [name, setName] = useState("");
  const [description, setDescription] = useState("");
  const [rules, setRules] = useState<z.infer<typeof BlockRulesSchema>>({
    nested: true,
    hasLikes: false,
    hasComments: false,
    commentsHasLikes: false,
    commentsEditable: false,
    commentsMaxNested: 0,
  });

  const handleSubmit = (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    mutate({
      name: name,
      description: description,
      rules: rules,
      author: "9b721bf6-4907-4467-a2da-d97feae31c9a",
      blocktype: "Post",
    });
    setOpen(false);
  };

  const queryClient = useQueryClient();
  const { mutate } = useMutation({
    mutationFn: (data: any) =>
      axios.post("cmd/block", data).then((res) => res.data),
    onSuccess: (response) => {
      console.log(response);
      queryClient.setQueryData(["posts"], (old: any) => [
        ...old,
        {
          name: name,
          description: description,
          id: response.id,
        },
      ]);
      queryClient.setQueryData(["block-meta", response.id], () => ({
        name: name,
        description: description,
      }));
      queryClient.setQueryData(["block-rules", response.id], () => ({
        ...rules,
      }));
    },
  });

  return (
    <form onSubmit={handleSubmit}>
      <div className="flex flex-col">
        <label htmlFor="name">Name</label>
        <input
          type="text"
          name="name"
          id="name"
          value={name}
          className="w-full border border-gray-200 rounded shadow-sm focus:outline-none focus:border-s-grass p-1"
          onChange={(e) => setName(e.currentTarget.value)}
        />
      </div>
      <div className="flex flex-col">
        <label htmlFor="description">Description</label>
        <textarea
          name="description"
          className="w-full border border-gray-200 rounded shadow-sm focus:outline-none focus:border-s-grass p-1"
          id="description"
          value={description}
          onChange={(e) => setDescription(e.currentTarget.value)}
        />
      </div>
      <div className="flex flex-col">
        <label htmlFor="nested">Nested</label>
        <input
          type="checkbox"
          name="nested"
          id="nested"
          checked={rules.nested}
          onChange={(e) =>
            setRules({ ...rules, nested: e.currentTarget.checked })
          }
        />
      </div>
      <div className="flex flex-col">
        <label htmlFor="hasLikes">hasLikes</label>
        <input
          type="checkbox"
          name="hasLikes"
          id="hasLikes"
          checked={rules.hasLikes}
          onChange={(e) =>
            setRules({ ...rules, hasLikes: e.currentTarget.checked })
          }
        />
      </div>
      <div className="flex flex-col">
        <label htmlFor="hasComments">hasComments</label>
        <input
          type="checkbox"
          name="hasComments"
          id="hasComments"
          checked={rules.hasComments}
          onChange={(e) =>
            setRules({ ...rules, hasComments: e.currentTarget.checked })
          }
        />
      </div>
      <div className="flex flex-col">
        <label htmlFor="commentsHasLikes">commentsHasLikes</label>
        <input
          type="checkbox"
          name="commentsHasLikes"
          id="commentsHasLikes"
          checked={rules.commentsHasLikes}
          onChange={(e) =>
            setRules({ ...rules, commentsHasLikes: e.currentTarget.checked })
          }
        />
      </div>
      <div className="flex flex-col">
        <label htmlFor="commentsEditable">commentsEditable</label>
        <input
          type="checkbox"
          name="commentsEditable"
          id="commentsEditable"
          checked={rules.commentsEditable}
          onChange={(e) =>
            setRules({ ...rules, commentsEditable: e.currentTarget.checked })
          }
        />
      </div>
      <div className="flex flex-col">
        <label htmlFor="commentsMaxNested">commentsMaxNested</label>
        <input
          type="number"
          name="commentsMaxNested"
          id="commentsMaxNested"
          value={rules.commentsMaxNested}
          onChange={(e) =>
            setRules({
              ...rules,
              commentsMaxNested: parseInt(e.currentTarget.value),
            })
          }
        />
      </div>
      <button type="submit">Submit</button>
    </form>
  );
}
