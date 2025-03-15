import classNames from "classnames";
import { useState } from "react";
import { ArrowUp, Paperclip } from "react-bootstrap-icons";

import { Button } from "../button/button";
import { TextArea } from "../textarea";

interface IProps {
  onSendMessage: (value: string) => void;
  onChange: (value: string) => void;
  onGrow?: (growing: boolean) => void;
  onKeyDown: (ev: React.KeyboardEvent<HTMLTextAreaElement>) => void;
  message: string;
}

export const SmallMessagebox = ({
  onSendMessage,
  onChange,
  onKeyDown,
  message,
  onGrow,
}: IProps) => {
  const [shouldGrow, setShouldGrow] = useState(false);

  const handleTextareaGrowth = (growing: boolean) => {
    onGrow?.(growing);
    setShouldGrow(growing);
  };

  const renderActionToolbar = () => {
    return (
      <div
        className={classNames("flex", {
          "justify-end w-full": shouldGrow,
        })}
      >
        <Button
          type="button"
          size="medium"
          variant="text"
          shape="circle"
          colorTheme="white"
        >
          <Paperclip size={20} className="rotate-45" />
        </Button>
        <Button
          type="button"
          size="medium"
          variant="contained"
          shape="circle"
          disabled={message.length === 0}
          onClick={() => message.length > 0 && onSendMessage(message)}
        >
          <ArrowUp size={20} />
        </Button>
      </div>
    );
  };

  return (
    <div
      className={classNames(
        "w-full relative rounded-full border shadow-sm p-2 focus-within:ring-1 border-gray-300 ring-gray-300 dark:border-white/15 dark:ring-white/15",
        { "rounded-md p-3": shouldGrow }
      )}
    >
      <div
        className={classNames("flex items-center", {
          "flex-col": shouldGrow,
        })}
      >
        <div
          className={classNames("w-full p-2 leading-[normal]", {
            "p-0": shouldGrow,
          })}
        >
          <TextArea
            placeholder="Ask follow-up"
            onChange={(ev: React.ChangeEvent<HTMLTextAreaElement>) => {
              if (ev.target.value.length > 45) {
                setShouldGrow(true);
              }

              onChange(ev.target.value.trim());
            }}
            onKeyDown={onKeyDown}
            onGrow={handleTextareaGrowth}
          />
        </div>
        {renderActionToolbar()}
      </div>
    </div>
  );
};
