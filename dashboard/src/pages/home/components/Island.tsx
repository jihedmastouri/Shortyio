import { HTMLAttributes, forwardRef } from "react";

const Island = forwardRef<HTMLDivElement, HTMLAttributes<HTMLDivElement>>(
	({ className, children }, ref) => (
		<div
			ref={ref}
			className={
				"p-3 my-2 w-full bg-white rounded-lg shadow dark:shadow-black dark:bg-gray-600 " +
				" " +
				className
			}
		>
			{children}
		</div>
	)
);

export default Island;
