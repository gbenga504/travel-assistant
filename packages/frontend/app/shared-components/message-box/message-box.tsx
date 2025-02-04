import { useState } from "react";

import { LargeMessagebox } from "./large-message-box";
import { SmallMessagebox } from "./small-message-box";

interface IProps {
  size: "small" | "large";
  onSendMessage: (value: string) => void;
}

export const Messagebox = ({ size, onSendMessage }: IProps) => {
  const [message, setMessage] = useState("");

  const handleKeyDown = (ev: React.KeyboardEvent<HTMLTextAreaElement>) => {
    if (ev.key === "Enter" || ev.keyCode === 13) {
      if (!ev.shiftKey && message.length > 0) {
        // We don't want the enter key to create a new line
        ev.preventDefault();

        onSendMessage(message);
      }
    }
  };

  const handleChange = (value: string) => {
    setMessage(value);
  };

  if (size === "large") {
    return (
      <LargeMessagebox
        onSendMessage={onSendMessage}
        onChange={handleChange}
        onKeyDown={handleKeyDown}
        message={message}
      />
    );
  }

  return (
    <SmallMessagebox
      onSendMessage={onSendMessage}
      onChange={handleChange}
      onKeyDown={handleKeyDown}
      message={message}
    />
  );
};
