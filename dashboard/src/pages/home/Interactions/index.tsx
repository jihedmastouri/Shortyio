import Header from "../components/Header";
import Comments from "./Comments";
import Likes from "./Likes";

const index = () => {
  return (
    <div className="w-full">
      <div className="mt-2 mx-5 flex justify-between">
        <Header />
      </div>
      {
        // <Tabs tabNames={["Likes", "Comments"]} tabs={[<Likes />, <Comments />]} />
      }
    </div>
  );
};

export default index;
