import classNames from "classnames";
import { GeoAltFill } from "react-bootstrap-icons";

import { useMapConfig } from "~/context/map-config-context";

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
  dataLongitude?: string;
  dataLatitude?: string;
}

const Span = ({ children, dataType, ...rest }: ISpanProps) => {
  const { changeMapConfig } = useMapConfig();

  const handleShowOnMap = () => {
    const { dataLongitude, dataLatitude, dataValue } = rest;

    if (dataLongitude && dataLatitude) {
      changeMapConfig({
        center: [Number(dataLatitude), Number(dataLongitude)],
        zoom: 7,
        markers: [
          {
            position: [Number(dataLatitude), Number(dataLongitude)],
            name: String(dataValue),
          },
        ],
      });
    }
  };

  switch (dataType) {
    case "location":
      return (
        <a
          {...rest}
          onFocus={handleShowOnMap}
          onMouseOver={handleShowOnMap}
          className={classNames(
            "inline-flex items-center font-medium cursor-pointer align-text-bottom h-5"
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
