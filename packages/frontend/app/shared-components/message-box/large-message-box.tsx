import { ArrowRight, PlusCircle } from "react-bootstrap-icons";

import { Button } from "../button/button";
import { TextArea } from "../Textarea";

interface IProps {
  onSendMessage: (value: string) => void;
  onChange: (value: string) => void;
  onKeyDown: (ev: React.KeyboardEvent<HTMLTextAreaElement>) => void;
  message: string;
}

export const LargeMessagebox = ({
  onSendMessage,
  onChange,
  onKeyDown,
  message,
}: IProps) => {
  const renderTextarea = () => {
    return (
      <div className="col-start-1 col-end-4">
        <TextArea
          onChange={(ev: React.ChangeEvent<HTMLTextAreaElement>) => {
            onChange(ev.target.value.trim());
          }}
          onKeyDown={onKeyDown}
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
            colorTheme="white"
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
    <div className="w-full relative rounded-md border shadow-sm p-4 pb-2 focus-within:ring-1 border-gray-300 ring-gray-300 dark:border-white/15 dark:ring-white/15">
      <div className="grid grid-rows-[1fr_auto] grid-cols-3">
        {renderTextarea()}
        {renderActionToolbar()}
      </div>
    </div>
  );
};
