import classNames from "classnames";
import { ReactNode } from "react";

interface IProps {
  children: ReactNode;
  className?: string;
}

export const MaxWidthContainer = ({ children, className }: IProps) => {
  return (
    <div className={classNames("max-w-screen-xl mx-auto", className)}>
      {children}
    </div>
  );
};
