import { ReactNode } from "react";
import * as RadSelect from "@radix-ui/react-select";
import {
	CheckIcon,
	ChevronDownIcon,
	ChevronUpIcon,
} from "@radix-ui/react-icons";

type itemType = {
	name: string | number;
	content: ReactNode;
};

type Props = {
	selected: string | number;
	ariaLabel: string;
	className?: string;
	items: itemType[];
	onChange: (value: string) => void;
};

const Select = ({ className, selected, items, onChange, ariaLabel }: Props) => {
	return (
		<RadSelect.Root defaultValue="Select an Option" value={String(selected)} onValueChange={onChange}>
			<RadSelect.Trigger
				className={
					"flex items-center justify-center text-sm outline-s-gopher" +
					" " +
					className
				}
				aria-label={ariaLabel}
			>
				<RadSelect.Value placeholder={ariaLabel + "..."} />
				<RadSelect.Icon className="w-5">
					<ChevronDownIcon />
				</RadSelect.Icon>
			</RadSelect.Trigger>
			<RadSelect.Portal>
				<RadSelect.Content
					className="overflow-hidden z-20 bg-white rounded-md max-h-[50vh] shadow-gray-300 shadow"
					position="popper"
					side="bottom"
					sideOffset={5}
				>
					<RadSelect.ScrollUpButton className="flex items-center justify-center h-5 bg-white cursor-default">
						<ChevronUpIcon />
					</RadSelect.ScrollUpButton>
					<RadSelect.Viewport className="p-1">
						{items.map((item, index) => (
							<RadSelect.Item
								key={index}
								value={"" + item.name}
								className="text-small leading-none rounded-[3px] flex items-center my-2 pr-4 pl-5 py-1 relative select-none cursor-pointer hover:bg-s-gopher-l border-none outline-s-gopher"
							>
								<RadSelect.ItemText>{item.content}</RadSelect.ItemText>
								<RadSelect.ItemIndicator className="absolute left-0 w-6 inline-flex items-center justify-center">
									<CheckIcon />
								</RadSelect.ItemIndicator>
							</RadSelect.Item>
						))}
					</RadSelect.Viewport>
					<RadSelect.ScrollDownButton className="flex items-center justify-center h-5 bg-white cursor-default">
						<ChevronDownIcon />
					</RadSelect.ScrollDownButton>
					<RadSelect.Arrow />
				</RadSelect.Content>
			</RadSelect.Portal>
		</RadSelect.Root>
	);
};

export default Select;
