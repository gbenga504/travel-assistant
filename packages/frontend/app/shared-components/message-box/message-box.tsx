import { ArrowRight, PlusCircle } from "react-bootstrap-icons";

import { Button } from "../button/button";
import { useEffect, useRef } from "react";

export const Messagebox = () => {
  const textareaRef = useRef<HTMLTextAreaElement>(null);

  useEffect(() => {
    const textarea = textareaRef.current;

    if (!textarea) {
      return;
    }

    function handleInput() {
      if (!textarea) return;

      // Reset height to auto to correctly calculate the scrollHeight
      textarea.style.height = "auto";

      // Calculate the new height on the content
      const newHeight = textarea.scrollHeight;

      // If the new height exceeds 300px, set it to 300px and make it scrollable
      if (newHeight > 200) {
        textarea.style.height = `200px`;
        textarea.style.overflowY = "scroll";
      } else {
        // Otherwise, adjust to fit content and hide scrolling
        textarea.style.height = `${newHeight}px`;
        textarea.style.overflowY = "hidden";
      }
    }

    textarea.addEventListener("input", handleInput);

    return () => {
      textarea.removeEventListener("input", handleInput);
    };
  }, []);

  const renderTextarea = () => {
    return (
      <div className="col-start-1 col-end-4">
        <textarea
          className="resize-none outline-none w-full font-light placeholder:text-gray-400 overflow-y-hidden"
          placeholder="Tell me about your trip..."
          ref={textareaRef}
        />
      </div>
    );
  };

  const renderActionToolbar = () => {
    return (
      <>
        <div className="self-center">
          <Button
            type="button"
            size="small"
            variant="text"
            shape="rounded"
            icon={<PlusCircle />}
          >
            Attach
          </Button>
        </div>
        <div className="col-start-3 col-end-4 self-center justify-self-end">
          <Button
            type="button"
            size="medium"
            variant="contained"
            shape="circle"
          >
            <ArrowRight size={20} />
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
