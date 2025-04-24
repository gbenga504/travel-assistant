import { SendFill } from "react-bootstrap-icons";

import { useUserSettings } from "~/context/user-settings-context";
import { Button } from "~/shared-components/button/button";

export const AppHeader = () => {
  const { userSettings } = useUserSettings();

  const renderLeftSection = () => {
    return <span className="text-sm font-medium dark:text-gray-300" />;
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
          {userSettings.userName[0].toUpperCase()}
        </div>
      </div>
    );
  };

  return (
    <div className="sticky top-0 w-full h-14 p-3 flex justify-between items-center">
      {renderLeftSection()}
      {renderRightSection()}
    </div>
  );
};
