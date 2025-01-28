import { PlusCircle } from "react-bootstrap-icons";

import { Button } from "../button/button";

export const Messagebox = () => {
  const renderTextarea = () => {
    return (
      <div className="col-start-1 col-end-4">
        <textarea
          className="resize-none outline-none w-full font-light placeholder:text-gray-400 "
          placeholder="Tell me about your trip..."
        />
      </div>
    );
  };

  const renderActionToolbar = () => {
    return (
      <>
        <div>
          <Button
            type="button"
            size="small"
            variant="text"
            rounded
            icon={<PlusCircle />}
          >
            Attach
          </Button>
        </div>
      </>
    );
  };

  return (
    <div className="w-full relative rounded-md border border-gray-300 shadow-sm p-4 pb-2 focus-within:ring-1 ring-gray-300">
      <div className="grid grid-rows-[1fr_auto] grid-cols-3">
        {renderTextarea()}
        {renderActionToolbar()}
      </div>
    </div>
  );
};
