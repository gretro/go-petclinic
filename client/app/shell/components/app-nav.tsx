import { NavLink } from "@remix-run/react";
import { useState } from "react";

type Menu = {
  id: string;
  text: string;
  icon?: string;
  to: string;
  exact?: boolean;
};

const MAIN_MENU_ITEMS: Menu[] = [
  {
    id: "dashboard",
    text: "Dashboard",
    to: "/",
    exact: true,
  },
  {
    id: "owners",
    text: "Owners & Pets",
    to: "/owners",
  },
  {
    id: "visits",
    text: "Visits",
    to: "/visits",
  },
];

export const AppNav = () => {
  const [isOpen, setOpen] = useState(false);

  const handleMenuToggle = () => {
    setOpen((open) => !open);
  };

  return (
    <nav className="bg-green-700 absolute top-0 left-0 right-0 z-50">
      <div className="px-2 sm:px-6 lg:px-8">
        <div className="relative flex h-16 items-center justify-between">
          {/* Mobile menu */}
          <div className="absolute inset-y-0 left-0 flex items-center sm:hidden">
            <button
              type="button"
              className="
                inline-flex 
                items-center 
                justify-center 
                rounded-md 
                p-2 
                text-gray-300 
                hover:bg-green-800 
                hover:text-white 
                focus:outline-none 
                focus:ring-2 
                focus:ring-inset 
                focus:ring-white"
              aria-controls="mobile-menu"
              aria-expanded={isOpen}
              onClick={handleMenuToggle}
            >
              {isOpen ? (
                <>
                  <span className="sr-only">Close menu</span>
                  <svg
                    className="h-6 w-6"
                    xmlns="http://www.w3.org/2000/svg"
                    fill="none"
                    viewBox="0 0 24 24"
                    strokeWidth="1.5"
                    stroke="currentColor"
                    aria-hidden="true"
                  >
                    <path strokeLinecap="round" strokeLinejoin="round" d="M6 18L18 6M6 6l12 12" />
                  </svg>
                </>
              ) : (
                <>
                  <span className="sr-only">Open menu</span>
                  <svg
                    className="block h-6 w-6"
                    xmlns="http://www.w3.org/2000/svg"
                    fill="none"
                    viewBox="0 0 24 24"
                    strokeWidth="1.5"
                    stroke="currentColor"
                    aria-hidden="true"
                  >
                    <path
                      strokeLinecap="round"
                      strokeLinejoin="round"
                      d="M3.75 6.75h16.5M3.75 12h16.5m-16.5 5.25h16.5"
                    />
                  </svg>
                </>
              )}
            </button>
          </div>

          <div className="flex flex-1 items-center justify-center sm:items-stretch sm:justify-start">
            <NavLink className="flex flex-shrink-0 items-center" to={"/"}>
              <h1 className="text-md sm:text-2xl text-white flex items-center gap-1">
                <span className="material-symbols-outlined">pets</span>
                Pet Clinic
              </h1>
            </NavLink>
            <div className="hidden sm:ml-6 sm:block">
              <div className="flex space-x-4">
                {MAIN_MENU_ITEMS.map((menuItem) => (
                  <NavLink
                    key={menuItem.id}
                    className="text-gray-300 hover:bg-green-800 hover:text-white px-3 py-2 rounded-md font-medium active:bg-black"
                    to={menuItem.to}
                  >
                    {menuItem.text}
                  </NavLink>
                ))}
              </div>
            </div>
          </div>
        </div>
      </div>

      {isOpen && (
        <div className="sm:hidden bg-green-700 absolute top-16 left-0 right-0">
          <div className="space-y-1 px-2 pt-2 pb-3">
            {MAIN_MENU_ITEMS.map((menuItem) => (
              <NavLink
                key={menuItem.id}
                className="text-gray-300 hover:bg-green-800 hover:text-white block px-3 py-2 rounded-md text-base font-medium active:bg-black"
                to={menuItem.to}
                onClick={handleMenuToggle}
              >
                {menuItem.text}
              </NavLink>
            ))}
          </div>
        </div>
      )}
    </nav>
  );
};
