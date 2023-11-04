import Editor from "./Editor";
import PageTabs from "@ui/TabsBackLink";
import Settings from "./Settings";
import { Navigate, useParams } from "react-router-dom";

type Props = {};

export default function index({ }: Props) {
  const id = useParams().id;

  if (!id) return <Navigate to="/posts" />;

	return (
		<div className="w-11/12 my-5">
			<PageTabs
				tabNames={["editPost", "settings", "Stats & Interactions"]}
				backLink={{
					name: "all Posts",
					link: "/posts",
				}}
				tabs={[
					<Editor  id={id} />,
					<Settings id={id} />,
          <div className="flex flex-col justify-center items-center"> hi</div>
				]}
			/>
		</div>
	);
}
