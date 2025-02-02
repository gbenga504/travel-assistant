import { Outlet } from "@remix-run/react";

import { MaxWidthContainer } from "~/shared-components/max-width-container";

export default function Route() {
  const renderSidebar = () => {
    return (
      <div className="fixed bg-white border-r border-r-gray-300 h-screen top-0 left-0 md:w-60"></div>
    );
  };

  return (
    <>
      {renderSidebar()}
      <MaxWidthContainer className="w-4/5 lg:w-[640px]">
        <Outlet />
      </MaxWidthContainer>
    </>
  );
}
