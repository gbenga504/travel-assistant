import classNames from "classnames";
import { GeoAltFill } from "react-bootstrap-icons";

import type { MarkdownToJSX } from "markdown-to-jsx";
import type { ReactNode } from "react";
import "./markdown-overrides.css";

const Li = ({ children, ...props }: { children: ReactNode }) => (
  <li {...props} className="mt-2">
    <span className="mr-1 markdown--li">- {children}</span>
  </li>
);

interface ISpanProps {
  children: ReactNode;
  dataType?: "userName" | "location" | "budget" | "travelDates";
  dataValue?: string;
  dataPreference?: "preferred";
}

const Span = ({ children, dataType, ...rest }: ISpanProps) => {
  switch (dataType) {
    case "location":
      return (
        <a
          {...rest}
          className={classNames(
            "inline-flex items-center font-medium cursor-pointer"
          )}
        >
          <GeoAltFill className="inline-block mr-1" />
          {children}
        </a>
      );
    default:
      return (
        <span {...rest} className={classNames("inline-flex", "items-center")}>
          {children}
        </span>
      );
  }
};

export const markdownOverrides: MarkdownToJSX.Overrides = {
  li: {
    component: Li,
  },
  span: {
    component: Span,
  },
};
