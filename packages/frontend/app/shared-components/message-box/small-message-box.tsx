import { ArrowUp, Paperclip } from "react-bootstrap-icons";
import { Button } from "../button/button";
import { TextArea } from "../Textarea";
import { useState } from "react";
import classNames from "classnames";

interface IProps {
  onSendMessage: (value: string) => void;
  onChange: (value: string) => void;
  onKeyDown: (ev: React.KeyboardEvent<HTMLTextAreaElement>) => void;
  message: string;
}

export const SmallMessagebox = ({
  onSendMessage,
  onChange,
  onKeyDown,
  message,
}: IProps) => {
  const [shouldGrow, setShouldGrow] = useState(false);

  const handleChangeHeight = (height: number) => {
    // The minimum height of the textarea is 20px
    // Hence if the height of the textarea is 20, then shouldGrow = false
    if (height <= 20) {
      return setShouldGrow(false);
    }

    setShouldGrow(true);
  };

  const renderActionToolbar = () => {
    return (
      <div
        className={classNames("flex", {
          "justify-end w-full": shouldGrow,
        })}
      >
        <Button type="button" size="medium" variant="text" shape="circle">
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
        "w-full relative rounded-full border border-gray-300 shadow-sm p-2 focus-within:ring-1 ring-gray-300",
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
            onChangeHeight={handleChangeHeight}
          />
        </div>
        {renderActionToolbar()}
      </div>
    </div>
  );
};
