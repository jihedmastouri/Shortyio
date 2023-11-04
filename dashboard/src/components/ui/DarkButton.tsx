import useDarkMode from "@hooks/darkMode";
import { MoonIcon, SunIcon } from "@radix-ui/react-icons";

const DarkButton = () => {
  // const [isDark, toggleMode] = useDarkMode();
  // onClick={() => toggleMode()}
  return (
    <button className="w-8 h-8 rounded-full font-bold flex mx-4 opacity-70 hover:bg-gray-200 hover:dark:bg-gray-800">
      {/* isDark ? (
				<MoonIcon className="text-gray-200" />
			) : (
				<SunIcon className="text-amber-600" />
			) */}
    </button>
  );
};

export default DarkButton;
