import DropDown from "@ui/DropDown";

const ChangeSite = () => {
  return (
    <div className="w-fit">
      <DropDown
        content={["New Tab", <div className="RightSlot">⌘+T</div>]}
        role="Change Domain"
      >
        <h2 className="font-bold text-2xl mx-2 py-2">MyBlog</h2>
      </DropDown>
    </div>
  );
};

export default ChangeSite;
