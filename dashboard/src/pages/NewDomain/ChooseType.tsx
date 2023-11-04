type WebsiteTypes = {
  id: number;
  name: string;
  img: string;
};

type props = {
  setSiteType: (params: number) => void;
  selected: number;
};

const types: WebsiteTypes[] = [
  {
    id: 0,
    name: "Blog",
    img: "https://via.placeholder.com/150.png",
  },
  {
    id: 1,
    name: "E-Commerce",
    img: "https://via.placeholder.com/450.png",
  },
];

const ChooseType = ({ setSiteType, selected }: props) => {
  const itemTypeClass = "mx-auto text-center p-5";
  return (
    <>
      <h1 className="text-center my-10">
        What Kind of website you cant to build?
      </h1>
      <div className="flex justify-around flex-wrap">
        {types.map((el) => {
          return (
            <div
              key={el.id}
              onClick={() => setSiteType(el.id)}
              className={
                (selected == el.id ? "bg-blue-50" : "") + " " + itemTypeClass
              }
            >
              <div className="w-[320px] h-[320px]  mx-auto flex justify-center">
                <img src={el.img} alt={el.name} />
              </div>
              <h2 className="mt-2">{el.name}</h2>
            </div>
          );
        })}
      </div>
    </>
  );
};

export default ChooseType;
