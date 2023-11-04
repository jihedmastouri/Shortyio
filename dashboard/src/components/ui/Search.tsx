import { MagnifyingGlassIcon } from "@radix-ui/react-icons";
import { useEffect, useState } from "react";

const Search = ({
	value: initialValue,
	onChange,
	debounce = 500,
	className,
	...props
}: {
	value: string | number;
	onChange: (value: string | number) => void;
	debounce?: number;
	className?: string;
} & Omit<
	React.InputHTMLAttributes<HTMLInputElement>,
	"onChange" | "className"
>) => {
	{
		const [value, setValue] = useState(initialValue);

		useEffect(() => {
			setValue(initialValue);
		}, [initialValue]);

		useEffect(() => {
			const timeout = setTimeout(() => {
				onChange(value);
			}, debounce);

			return () => clearTimeout(timeout);
		}, [value]);

		return (
			<div
				className={
					"flex bg-white outline-none focus-within:outline-s-gopher h-10" +
					" " +
					className
				}
			>
				<MagnifyingGlassIcon className="w-5 float-left" />
				<input
					{...props}
					className="h-full w-full outline-none"
					value={value}
					onChange={(e) => setValue(e.target.value)}
				/>
			</div>
		);
	}
};

export default Search;
