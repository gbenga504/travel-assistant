import { LargeMessagebox } from "./large-message-box";
import { SmallMessagebox } from "./small-message-box";

interface IProps {
  size: "small" | "large";
  onGrow?: (growing: boolean) => void;
  onSendMessage: (value: string) => void;
  value: string;
  onChange: (message: string) => void;
}

export const Messagebox = ({
  size,
  onSendMessage,
  onGrow,
  value,
  onChange,
}: IProps) => {
  const handleKeyDown = (ev: React.KeyboardEvent<HTMLTextAreaElement>) => {
    if (ev.key === "Enter" || ev.keyCode === 13) {
      if (!ev.shiftKey && value.length > 0) {
        // We don't want the enter key to create a new line
        ev.preventDefault();

        // If we have a value then it is a controlled component and we use that instead
        onSendMessage(value);
      } else if (ev.shiftKey) {
        onChange(`${value}`);
      }
    }
  };

  const handleChange = (value: string) => {
    onChange(value);
  };

  if (size === "large") {
    return (
      <LargeMessagebox
        onSendMessage={onSendMessage}
        onChange={handleChange}
        onKeyDown={handleKeyDown}
        value={value}
      />
    );
  }

  return (
    <SmallMessagebox
      onSendMessage={onSendMessage}
      onChange={handleChange}
      onKeyDown={handleKeyDown}
      value={value}
      onGrow={onGrow}
    />
  );
};
