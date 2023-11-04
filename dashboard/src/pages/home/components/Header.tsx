import DarkButton from "@ui/DarkButton";
import ChangeSite from "./ChangeSite";

const Header = () => {
  return (
    <div className="h-full flex items-center">
      <ChangeSite />
      <DarkButton />
    </div>
  );
};

export default Header;
