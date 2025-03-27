import classNames from "classnames";
import { useEffect, useRef } from "react";

interface IProps extends React.TextareaHTMLAttributes<HTMLTextAreaElement> {
  value: string;
  onGrow?: (growing: boolean) => void;
}

export const TextArea = ({ className, onGrow, value, ...rest }: IProps) => {
  const textareaRef = useRef<HTMLTextAreaElement>(null);

  useEffect(() => {
    const textarea = textareaRef.current;

    if (!textarea) {
      return;
    }

    // Reset height to auto to correctly calculate the scrollHeight
    textarea.style.height = "0px";

    // Calculate the new height on the content
    const newHeight = Math.min(textarea.scrollHeight, 200);

    if (newHeight >= 200) {
      textarea.style.overflowY = "scroll";
    } else {
      textarea.style.overflowY = "hidden";
    }

    textarea.style.height = `${newHeight}px`;

    // If the newHeight is more than 20px, then we want the textarea to grow
    // However, if the value of the textarea is more than 45 then we also want it to grow
    let shouldGrow = newHeight > 20 ? true : false;

    if (textarea.value.length > 45) {
      shouldGrow = true;
    }

    onGrow?.(shouldGrow);
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [value]);

  return (
    <textarea
      className={classNames(
        "resize-none outline-none w-full font-light overflow-y-hidden bg-inherit placeholder:text-gray-400",
        className
      )}
      placeholder="Tell me about your trip..."
      ref={textareaRef}
      value={value}
      rows={1}
      {...rest}
    />
  );
};
