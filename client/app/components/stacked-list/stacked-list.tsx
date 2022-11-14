import { Link } from "@remix-run/react";
import { useRef } from "react";

export type ColumnDef = {
  id: string;
  className?: string;
  name: string;
  fieldName: string;
};

export const StackedList = () => {
  const items = new Array(10).fill(null).map((_, i) => ({
    id: i,
    firstName: "René-Pier",
    lastName: "Deshaies Gélinas",
    phoneNumber: "514-555-2838",
    createdAt: new Date("2022-01-01"),
  }));

  const intl = useRef(new Intl.DateTimeFormat("en"));

  return (
    <ul className="bg-gray-100 dark:bg-gray-600 shadow-md border-y-2 border-gray-300 dark:border-gray-500 rounded-md dark:text-white">
      {items.map((item) => (
        <li key={item.id} className="flex px-4 py-4 first:border-t-0 border-t-2 dark:border-t-gray-500">
          <div className="flex-1 flex flex-col justify-center">
            <p className="font-semibold">
              {item.lastName}, {item.firstName}
            </p>
            <p className="flex items-center gap-1 font-light text-sm text-gray-500">
              <span className="material-symbols-outlined" style={{ fontSize: "1.125rem", lineHeight: "1.125rem" }}>
                call
              </span>
              {item.phoneNumber}
            </p>
          </div>
          <div className="hidden sm:flex flex-1 flex-col justify-center">
            <p>Client since</p>
            <p className="font-light text-sm text-gray-500">{intl.current.format(item.createdAt)}</p>
          </div>
          <div>
            <Link className="h-full flex items-center" to={String(item.id)}>
              <span className="material-symbols-outlined" style={{ fontSize: "2.25rem" }}>
                chevron_right
              </span>
            </Link>
          </div>
        </li>
      ))}
    </ul>
  );
};
