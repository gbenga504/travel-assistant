import classNames from "classnames";
import { useEffect, useRef } from "react";

interface IProps extends React.TextareaHTMLAttributes<HTMLTextAreaElement> {}

export const TextArea = ({ className, ...rest }: IProps) => {
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

  return (
    <textarea
      className={classNames(
        "resize-none outline-none w-full font-light placeholder:text-gray-400 overflow-y-hidden",
        className
      )}
      placeholder="Tell me about your trip..."
      ref={textareaRef}
      rows={1}
      {...rest}
    />
  );
};
