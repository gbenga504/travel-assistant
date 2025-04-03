import classNames from "classnames";

import type { MarkdownToJSX } from "markdown-to-jsx";
import type { ReactNode } from "react";

const Li = ({ children, ...props }: { children: ReactNode }) => (
  <li {...props} className="mt-2">
    - {children}
  </li>
);

interface ISpanProps {
  children: ReactNode;
  dataType?: "userName" | "location" | "budget" | "travelDates";
  dataValue?: string;
  dataPreference?: "preferred";
}

const Span = ({ children, dataType, ...rest }: ISpanProps) => (
  <span
    {...rest}
    className={classNames({
      "font-medium": dataType !== "userName",
    })}
  >
    {children}
  </span>
);

export const markdownOverrides: MarkdownToJSX.Overrides = {
  li: {
    component: Li,
  },
  span: {
    component: Span,
  },
};
