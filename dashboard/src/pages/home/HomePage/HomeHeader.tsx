import * as Avatar from "@radix-ui/react-avatar";
import DropDown from "@ui/DropDown";
import Header from "../components/Header";

type info = {
  image: string;
  name: string;
  lastName: string;
};

const HomeHeader = () => {
  const user: info = {
    image:
      "https://images.unsplash.com/photo-1492633423870-43d1cd2775eb?&w=128&h=128&dpr=2&q=80",
    // image: "",
    name: "salah",
    lastName: "salah",
  };

  const username = (user: info) => {
    return user.name[0] + user.lastName[0];
  };

  return (
    <div className="dark:bg-gray-600 w-full h-20 flex shadow light:shadow-gray-100 rounded-tr-md">
      <div className="w-[98%] h-16 m-auto flex justify-between items-center">
        <Header />
        <DropDown
          content={["New Tab", <div className="RightSlot">⌘+T</div>]}
          role="Change Domain"
        >
          <div className="h-full flex items-center">
            <div className="">hiiiiiiiiiiiiiiii</div>
            <Avatar.Root className="block rounded-full aspect-square h-2/3 m-1">
              <Avatar.Image
                className="object-cover h-full rounded-full"
                src={user.image}
                alt="Colm Tuite"
              />
              <Avatar.Fallback
                className="w-full h-full rounded-full flex justify-center items-center bg-white text-s-grass shadow-gray shadow font-bold align-baseline"
                delayMs={600}
              >
                {username(user)}
              </Avatar.Fallback>
            </Avatar.Root>
          </div>
        </DropDown>
      </div>
    </div>
  );
};

export default HomeHeader;
