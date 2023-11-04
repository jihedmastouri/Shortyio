import Tabs from "@ui/Tabs";
import Header from "../components/Header";
import Customers from "./Customers";
import Editors from "./Editors";

const index = () => {
	return (
		<div className="w-full">
			<div className="mt-2 mx-5 flex justify-between">
				<Header />
			</div>
			<Tabs
				tabNames={["Editor/Admins", "Customers"]}
				tabs={[<Editors />, <Customers />]}
			/>
		</div>
	);
};

export default index;
