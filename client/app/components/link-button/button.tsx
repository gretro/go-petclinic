import { Link } from "@remix-run/react";
import type { PropsWithChildren } from "react";

type Props = {
  flavour?: "primary" | "secondary" | "default";
  to: string;
} & PropsWithChildren;

export const LinkButton = (props: Props) => {
  let btnClass = "bg-gray-500 hover:bg-gray-400 ring-gray-500 disabled:bg-gray-500";

  if (props.flavour === "primary") {
    btnClass = "bg-green-700 hover:bg-green-600 ring-green-700 disabled:bg-green-700";
  }

  return (
    <Link
      className={`${btnClass} px-4 py-2 ring-1 shadow-md rounded-md text-white disabled:opacity-60 disabled:ring-0 disabled:text-gray-400`}
      to={props.to}
    >
      {props.children}
    </Link>
  );
};
