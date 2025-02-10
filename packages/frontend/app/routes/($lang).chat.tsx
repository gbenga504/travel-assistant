import { Outlet } from "@remix-run/react";

export default function Route() {
  const renderSidebar = () => {
    return (
      <div className="fixed border-r h-screen top-0 left-0 md:w-60 bg-white border-gray-200 dark:bg-gray-900 dark:border-white/10"></div>
    );
  };

  return (
    <>
      {renderSidebar()}
      <Outlet />
    </>
  );
}
