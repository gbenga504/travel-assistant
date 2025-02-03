import { Outlet } from "@remix-run/react";

export default function Route() {
  const renderSidebar = () => {
    return (
      <div className="fixed bg-white border-r border-r-gray-300 h-screen top-0 left-0 md:w-60"></div>
    );
  };

  return (
    <>
      {renderSidebar()}
      <Outlet />
    </>
  );
}
