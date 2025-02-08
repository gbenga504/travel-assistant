import { useState, useRef, useEffect, cloneElement, ReactElement } from "react";
import { Button } from "./button/button";

interface IOption {
  value: string;
  label: string;
}

interface IProps {
  trigger: ReactElement;
  options: IOption[];
  onSelect: (option: IOption) => void;
}

export const Dropdown = ({ trigger, options, onSelect }: IProps) => {
  const [isOpen, setIsOpen] = useState(false);
  const dropdownRef = useRef<HTMLDivElement>(null);

  useEffect(() => {
    const handleClickOutside = (event: MouseEvent) => {
      if (
        dropdownRef.current &&
        !dropdownRef.current.contains(event.target as Node)
      ) {
        setIsOpen(false);
      }
    };

    document.addEventListener("mousedown", handleClickOutside);

    return () => {
      document.removeEventListener("mousedown", handleClickOutside);
    };
  }, []);

  const handleToggleDropdown = () => setIsOpen(!isOpen);

  const handleOptionClick = (option: IOption) => {
    setIsOpen(false);
    onSelect(option);
  };

  const handleKeyDown = (option: IOption) => {
    return (ev: React.KeyboardEvent<HTMLButtonElement>) => {
      if (ev.key === "Enter" || ev.keyCode === 13) {
        handleOptionClick(option);
      }
    };
  };

  return (
    <div className="relative" ref={dropdownRef}>
      {cloneElement(trigger, { onClick: handleToggleDropdown })}

      {isOpen && (
        <ul className="absolute top-full left-0 z-50 min-w-40 list-none bg-white rounded-md shadow-md my-1">
          {options.map((option) => (
            <li key={option.value} className="rounded-md">
              <Button
                type="button"
                variant="text"
                size="large"
                colorTheme="white"
                onClick={() => handleOptionClick(option)}
                onKeyDown={handleKeyDown(option)}
                fullWidth
                className="!text-base"
              >
                {option.label}
              </Button>
            </li>
          ))}
        </ul>
      )}
    </div>
  );
};
