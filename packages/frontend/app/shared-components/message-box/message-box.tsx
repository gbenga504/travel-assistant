import { useState } from "react";
import { ArrowRight, PlusCircle } from "react-bootstrap-icons";

import { Button } from "../button/button";
import { TextArea } from "../Textarea";

interface IProps {
  onSendMessage: (value: string) => void;
}

export const Messagebox = ({ onSendMessage }: IProps) => {
  const [message, setMessage] = useState("");

  const handleKeyUp = (ev: React.KeyboardEvent<HTMLTextAreaElement>) => {
    if (
      (ev.key === "Enter" || ev.keyCode === 13) &&
      !ev.shiftKey &&
      message.length > 0
    ) {
      console.log("gad ==> message ", message.replace(/\n/g, ""));
      onSendMessage(message);
    }
  };

  const renderTextarea = () => {
    return (
      <div className="col-start-1 col-end-4">
        <TextArea
          onChange={(ev: React.ChangeEvent<HTMLTextAreaElement>) => {
            setMessage(ev.target.value.trim());
          }}
          onKeyUp={handleKeyUp}
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
            disabled={message.length === 0}
            onClick={() => message.length > 0 && onSendMessage(message)}
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
