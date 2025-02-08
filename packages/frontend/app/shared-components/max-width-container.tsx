import classNames from "classnames";
import { ReactNode } from "react";

interface IProps {
  children: ReactNode;
  className?: string;
}

export const MaxWidthContainer = ({ children, className }: IProps) => {
  return (
    <div
      className={classNames(
        "px-5 md:px-16 xl:px-0 xl:max-w-screen-lg 2xl:max-w-screen-xl mx-auto",
        className
      )}
    >
      {children}
    </div>
  );
};
