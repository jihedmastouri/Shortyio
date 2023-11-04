import * as Switch from "@radix-ui/react-switch";
import { useState } from "react";

type Props = {
  id: string;
  value: boolean;
  onChange: (id: string, value: boolean) => void;
};

const SwitchToggle = ({ onChange, id, value }: Props) => {
  const [checked, setChecked] = useState(value);

  const handleChange = (v: boolean) => {
    setChecked(v);
    onChange(id, v);
  };

  return (
    <Switch.Root
      className="rounded-full w-10 h-5 bg-s-grass-l data-[state=checked]:bg-s-gopher shadow-md"
      id={id}
      name={id}
      checked={checked}
      onCheckedChange={handleChange}
    >
      <Switch.Thumb className="block w-5 h-5 rounded-full bg-white transition-transform duration-100 translate-x-0.5 will-change-transform data-[state=checked]:translate-x-5" />
    </Switch.Root>
  );
};

export default SwitchToggle;
