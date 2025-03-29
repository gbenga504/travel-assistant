import { SendFill } from "react-bootstrap-icons";

import { Button } from "~/shared-components/button/button";

export const AppHeader = () => {
  const renderLeftSection = () => {
    return <span className="text-sm font-medium dark:text-gray-300" />;
  };

  const renderMiddleSection = () => {
    return (
      <div className="py-2 flex text-sm border border-gray-200 rounded-xl text-gray-900 dark:text-gray-300 dark:bg-gray-900 dark:border-white/10">
        <div className="border-r px-3 border-r-gray-200 dark:border-white/10">
          <span>Prague</span>
        </div>
        <div className="border-r px-3 border-r-gray-200 dark:border-white/10">
          <span>June 18 - 20</span>
        </div>
        <div className="border-r px-3 border-r-gray-200 dark:border-white/10">
          <span>2 travelers</span>
        </div>
        <div className="px-3">
          <span>$</span>
        </div>
      </div>
    );
  };

  const renderRightSection = () => {
    return (
      <div className="flex items-center gap-x-2">
        <Button
          type="button"
          variant="contained"
          size="small"
          colorTheme="white"
          icon={<SendFill />}
        >
          Share
        </Button>
        <div className="h-8 w-8 font-bold rounded-full flex justify-center items-center bg-gray-200 dark:bg-gray-900 text-sm">
          G
        </div>
      </div>
    );
  };

  return (
    <div className="sticky top-0 w-full h-14 p-3 flex justify-between items-center">
      {renderLeftSection()}
      {renderMiddleSection()}
      {renderRightSection()}
    </div>
  );
};
