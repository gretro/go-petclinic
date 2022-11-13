import type { PropsWithChildren } from "react";

type Props = {
  title: string;
  subTitle?: string;
} & PropsWithChildren;

export const PageHeader = (props: Props) => {
  return (
    <header className="bg-gray-100 dark:bg-gray-600 shadow -mx-8 -mt-4 mb-4">
      <div className="mx-auto py-6 px-4 sm:px-6 lg:px-8 flex items-center justify-between gap-2">
        <div className="flex-1">
          <h2 className="text-xl sm:text-3xl font-bold tracking-tight text-gray-900 dark:text-gray-100">
            {props.title}
          </h2>
          {props.subTitle && <p className="text-sm text-gray-500 dark:text-gray-300">{props.subTitle}</p>}
        </div>
        {props.children}
      </div>
    </header>
  );
};
