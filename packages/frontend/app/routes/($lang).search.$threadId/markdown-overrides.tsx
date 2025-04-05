import classNames from "classnames";
import { GeoAltFill } from "react-bootstrap-icons";

import type { MarkdownToJSX } from "markdown-to-jsx";
import type { ReactNode } from "react";

const Li = ({ children, ...props }: { children: ReactNode }) => (
  <li {...props} className="mt-2">
    <span className="mr-1">- {children}</span>
  </li>
);

interface ISpanProps {
  children: ReactNode;
  dataType?: "userName" | "location" | "budget" | "travelDates";
  dataValue?: string;
  dataPreference?: "preferred";
}

const Span = ({ children, dataType, ...rest }: ISpanProps) => {
  return (
    <span
      {...rest}
      className={classNames("inline-flex", "items-center", {
        "font-medium": dataType === "location",
      })}
    >
      {dataType === "location" && <GeoAltFill className="inline-block mr-1" />}
      {children}
    </span>
  );
};

export const markdownOverrides: MarkdownToJSX.Overrides = {
  li: {
    component: Li,
  },
  span: {
    component: Span,
  },
};
