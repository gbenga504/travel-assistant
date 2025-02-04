import { ArrowUp, Paperclip } from "react-bootstrap-icons";
import { Button } from "../button/button";
import { TextArea } from "../Textarea";

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
  const renderActionToolbar = () => {
    return (
      <div className="flex">
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
    <div className="w-full relative rounded-full border border-gray-300 shadow-sm p-2 focus-within:ring-1 ring-gray-300">
      <div className="flex items-center">
        <div className="w-full p-2 leading-[normal]">
          <TextArea
            placeholder="Ask follow-up"
            onChange={(ev: React.ChangeEvent<HTMLTextAreaElement>) => {
              onChange(ev.target.value.trim());
            }}
            onKeyDown={onKeyDown}
          />
        </div>
        <div>{renderActionToolbar()}</div>
      </div>
    </div>
  );
};
