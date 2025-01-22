import { NavLink } from "@remix-run/react";
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
    <nav className="border-b border-b-gray-200 fixed left-0 right-0 z-30">
      <MaxWidthContainer className="lg:min-h-20 p-4 flex justify-between items-center">
        <div />
        <div>
          <NavLink to="/pricing" className={getNavLinkClassNames}>
            Pricing
          </NavLink>
        </div>
      </MaxWidthContainer>
    </nav>
  );
};
