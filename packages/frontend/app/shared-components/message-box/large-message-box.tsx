import { ArrowRight, PlusCircle } from "react-bootstrap-icons";

import { Button } from "../button/button";
import { TextArea } from "../textarea";

interface IProps {
  onSendMessage: (value: string) => void;
  onChange: (value: string) => void;
  onKeyDown: (ev: React.KeyboardEvent<HTMLTextAreaElement>) => void;
  value: string;
}

export const LargeMessagebox = ({
  onSendMessage,
  onChange,
  onKeyDown,
  value,
}: IProps) => {
  const renderTextarea = () => {
    return (
      <div className="col-start-1 col-end-4">
        <TextArea
          onChange={(ev: React.ChangeEvent<HTMLTextAreaElement>) => {
            onChange(ev.target.value);
          }}
          onKeyDown={onKeyDown}
          value={value}
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
            disabled={value.trim().length === 0}
            onClick={() => value.trim().length > 0 && onSendMessage(value)}
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
