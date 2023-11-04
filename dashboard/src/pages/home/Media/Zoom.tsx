import * as Dialog from "@radix-ui/react-dialog";
import { Cross2Icon } from "@radix-ui/react-icons";
import { HTMLAttributes } from "react";

const Zoom = ({
	img,
	alt,
	children,
	className,
}: { img: string; alt: string } & HTMLAttributes<HTMLDivElement>) => {
	return (
		<Dialog.Root>
			<Dialog.Trigger className={className}> {children} </Dialog.Trigger>
			<Dialog.Portal>
				<Dialog.Overlay className="z-40 bg-gray-600 fixed inset-0 opacity-30" />
				<Dialog.Content className="z-50 fixed top-[50%] left-[50%] translate-x-[-50%] translate-y-[-50%] flex flex-col">
					<img
						className="max-w-[90vw] max-h-[90vh] m-auto"
						src={img}
						alt={alt}
					/>
					<Dialog.Close asChild>
						<button
							className="w-5 h-5 absolute rounded bg-white self-end mt-[-1em] mr-[-2em] shadow"
							aria-label="Close"
						>
							<Cross2Icon />
						</button>
					</Dialog.Close>
				</Dialog.Content>
			</Dialog.Portal>
		</Dialog.Root>
	);
};

export default Zoom;
