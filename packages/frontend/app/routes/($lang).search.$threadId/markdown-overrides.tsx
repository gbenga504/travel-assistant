import classNames from "classnames";
import { Geo, Shop } from "react-bootstrap-icons";

import { useMapConfig } from "~/context/map-config-context";

import type { MarkdownToJSX } from "markdown-to-jsx";
import type { ReactNode } from "react";

const Li = ({ children, ...props }: { children: ReactNode }) => (
  <li {...props} className="mt-2">
    {children}
  </li>
);

interface ICoordinates {
  name: string;
  latitude: string;
  longitude: string;
  children: ReactNode;
}

const Location = ({ name, latitude, longitude }: ICoordinates) => {
  const { changeMapConfig } = useMapConfig();

  const handleShowOnMap = () => {
    if (longitude && latitude) {
      changeMapConfig({
        center: [Number(latitude), Number(longitude)],
        zoom: 7,
        markers: [
          {
            position: [Number(latitude), Number(longitude)],
            name: String(name),
          },
        ],
      });
    }
  };

  return (
    <span
      onFocus={handleShowOnMap}
      onMouseOver={handleShowOnMap}
      className={classNames(
        "inline-flex items-center font-medium cursor-pointer align-text-bottom h-5"
      )}
    >
      <Geo className="inline-block mr-1" />
      {name}
    </span>
  );
};

const Attraction = ({ name, latitude, longitude }: ICoordinates) => {
  const { changeMapConfig } = useMapConfig();

  const handleShowOnMap = () => {
    if (longitude && latitude) {
      changeMapConfig({
        center: [Number(latitude), Number(longitude)],
        zoom: 7,
        markers: [
          {
            position: [Number(latitude), Number(longitude)],
            name: String(name),
          },
        ],
      });
    }
  };

  return (
    <span
      onFocus={handleShowOnMap}
      onMouseOver={handleShowOnMap}
      className={classNames(
        "inline-flex items-center font-medium cursor-pointer align-text-bottom h-5"
      )}
    >
      <Shop className="inline-block mr-1" />
      {name}
    </span>
  );
};

export const markdownOverrides: MarkdownToJSX.Overrides = {
  li: {
    component: Li,
  },
  Location: {
    component: Location,
  },
  Attraction: {
    component: Attraction,
  },
};
