import { Outlet, useParams } from "@remix-run/react";
import { ReactElement, ReactNode } from "react";
import { Backpack, Wallet2 } from "react-bootstrap-icons";

import { Button } from "~/shared-components/button/button";
import { constructURL, ROUTE_IDS } from "~/utils/route-util";

interface ISidebarItem {
  icon: ReactElement;
  children: ReactNode;
  to: string;
}

const SidebarItem = ({ to, children, icon }: ISidebarItem) => {
  return (
    <Button
      type="link"
      variant="text"
      colorTheme="white"
      size="large"
      to={to}
      fullWidth
      icon={icon}
      className="!justify-start !text-base"
    >
      {children}
    </Button>
  );
};

export default function Route() {
  const { lang } = useParams();

  const renderSidebar = () => {
    return (
      <div className="fixed hidden flex-col border-r h-screen top-0 left-0 pt-4 px-1 w-60 bg-white border-gray-200 dark:bg-gray-900 dark:border-white/10 lg:flex">
        <div className="mb-4 w-full" />
        <div className="pt-4 w-full overflow-x-hidden">
          <SidebarItem
            to={constructURL({ routeId: ROUTE_IDS.homePage, params: { lang } })}
            icon={<Backpack className="mr-1" />}
          >
            Home
          </SidebarItem>
          <SidebarItem
            to={constructURL({
              routeId: ROUTE_IDS.pricingPage,
              params: { lang },
            })}
            icon={<Wallet2 className="mr-1" />}
          >
            Pricing
          </SidebarItem>
        </div>
        <div className="absolute bottom-0 full px-5 pb-5">
          <span className="font-light text-xs mb-3">
            &copy; WakaTravel {new Date().getFullYear()}
          </span>
        </div>
      </div>
    );
  };

  return (
    <>
      {renderSidebar()}
      <div className="ml-0 w-full h-full lg:ml-60 lg:w-[calc(100%-240px)]">
        <Outlet />
      </div>
    </>
  );
}
