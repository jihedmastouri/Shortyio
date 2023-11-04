import { useState } from "react";
import PrograssBar from "@ui/PrograssBar";
import ChooseType from "./ChooseType";

const CreateSite = () => {
  const [siteType, setSiteType] = useState(0);
  const progress = 66;
  const step = 0;

  const renderSwitch = (param: number) => {
    switch (param) {
      case 0:
        return <ChooseType selected={siteType} setSiteType={setSiteType} />;
      case 1:
        return "foo";
    }
  };

  return (
    <div className="h-screen flex flex-col">
      <div className="mx-auto md:my-10 my-1 h-5 w-1/2 flex flex-col">
        <PrograssBar progress={progress} />
      </div>
      <div className="w-[80%] mx-auto mt-10">{renderSwitch(step)}</div>
      <div className="flex justify-between w-[80%] mx-auto my-10">




















        <button>Previous</button>
        <button>{step < 3 ? "next" : "submit"}</button>
      </div>
    </div>
  );
};

export default CreateSite;
