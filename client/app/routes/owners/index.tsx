import { LinkButton } from "~/components/link-button/button";
import { StackedList } from "~/components/stacked-list/stacked-list";
import { PageHeader } from "~/shell/components/page-header";

const OwnersPage = () => {
  return (
    <>
      <PageHeader title="Owners & Pets Management" subTitle="Consult and manage the clinic's pet Owners and their Pets">
        <div>
          <LinkButton flavour="primary" to="new">
            Add Owner
          </LinkButton>
        </div>
      </PageHeader>

      <label className="relative block">
        <span className="sr-only">Search</span>
        <span className="absolute inset-y-0 left-0 flex items-center pl-2">
          <span className="material-symbols-outlined text-slate-800 dark:text-slate-200">search</span>
        </span>
        <input
          className="mb-4 rounded-md placeholder:italic placeholder:text-slate-400 block dark:text-white bg-white dark:bg-gray-600 w-full border border-slate-300 py-2 pl-9 pr-3 shadow-sm focus:outline-none focus:border-sky-500 focus:ring-sky-500 focus:ring-1 sm:text-sm"
          placeholder="Search owners"
        ></input>
      </label>
      <StackedList></StackedList>
    </>
  );
};

export default OwnersPage;
