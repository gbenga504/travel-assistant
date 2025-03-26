import type { MarkdownToJSX } from "markdown-to-jsx";
import type { ReactNode } from "react";

const Li = ({ children, ...props }: { children: ReactNode }) => (
  <li {...props} className="mt-2">
    - {children}
  </li>
);

export const markdownOverrides: MarkdownToJSX.Overrides = {
  li: {
    component: Li,
  },
};
