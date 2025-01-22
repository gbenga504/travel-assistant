import { NavLink } from "@remix-run/react";
import { Globe } from "react-bootstrap-icons";

import { MaxWidthContainer } from "../../shared-components/max-width-container";

export const AppHeader = () => {
  const getNavLinkClassNames = ({ isActive }: { isActive: boolean }) => {
    const defaultClasses = "py-2 px-3 font-light hover:text-blue-600";

    if (isActive) {
      return `${defaultClasses} text-blue-500`;
    }

    return defaultClasses;
  };

  return (
    <nav className="border-b border-b-gray-200 fixed left-0 top-0 z-30 w-full">
      <MaxWidthContainer className="lg:min-h-20 p-4 flex justify-between items-center">
        <div />
        <div className="flex">
          <NavLink to="/pricing" className={getNavLinkClassNames}>
            Pricing
          </NavLink>
          <NavLink to="/lang" className={getNavLinkClassNames}>
            <div className="flex items-center">
              <Globe className="mr-1" /> English
            </div>
          </NavLink>
        </div>
      </MaxWidthContainer>
    </nav>
  );
};
