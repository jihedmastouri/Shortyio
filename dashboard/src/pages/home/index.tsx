import { Outlet } from "react-router-dom";
import Header from "./components/Header";
import SideBar from "./components/SideBar";

function index() {
  return (
    <>
      <div className="h-full w-full flex">
        <SideBar />
        <div className="w-full h-full flex flex-col items-center overflow-y-scroll">
          <Outlet />
        </div>
      </div>
    </>
  );
}

export default index;
