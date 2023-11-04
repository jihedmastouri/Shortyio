import { useEffect } from "react";
import useLocalStorage from "./localStorage";

const useDarkMode = () => {
	const [enabled, setEnabled, isAvailable] =
		useLocalStorage<boolean>("dark-theme");

	if (
		window.matchMedia &&
		window.matchMedia("(prefers-color-scheme: dark)").matches &&
		!isAvailable
	) {
		setEnabled(true);
	}

	const toogleDark = () => {
		setEnabled((el: boolean) => !el);
	};

	useEffect(() => {
		const className = "dark";
		const bodyClass = window.document.body.classList;

		enabled ? bodyClass.add(className) : bodyClass.remove(className);
	}, [enabled]);

	return [enabled, toogleDark, setEnabled] as const;
};

export default useDarkMode;
