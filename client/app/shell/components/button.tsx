import type { PropsWithChildren, SyntheticEvent } from "react";

type Props = {
  flavour?: "primary" | "secondary" | "default";
  disabled?: boolean;
  onClick?: (e: SyntheticEvent) => void;
} & PropsWithChildren;

export const Button = (props: Props) => {
  let btnClass = "bg-gray-500 hover:bg-gray-400 ring-gray-500 disabled:bg-gray-500";

  if (props.flavour === "primary") {
    btnClass = "bg-green-700 hover:bg-green-600 ring-green-700 disabled:bg-green-700";
  }

  return (
    <button
      className={`${btnClass} px-4 py-2 ring-1 shadow-md rounded-md text-white disabled:opacity-60 disabled:ring-0 disabled:text-gray-400`}
      disabled={props.disabled}
      onClick={props.onClick}
    >
      {props.children}
    </button>
  );
};
