import { createContext, useContext, useState } from "react";

import type { LatLngTuple } from "leaflet";
import type React from "react";
import type { ReactNode } from "react";

interface IMapConfig {
  center: LatLngTuple;
  zoom: number;
  markers: { position: LatLngTuple; name: string }[];
}

const defaultMapConfig: IMapConfig = {
  center: [52.52437, 13.41053],
  zoom: 12,
  markers: [{ position: [52.52437, 13.41053], name: "Berlin" }],
};

const MapConfigContext = createContext<{
  mapConfig: IMapConfig;
  changeMapConfig: (mc: Partial<IMapConfig>) => void;
}>({
  mapConfig: defaultMapConfig,
  changeMapConfig: (mc: Partial<IMapConfig>) => mc,
});

export const useMapConfig = () => {
  return useContext(MapConfigContext);
};

export const MapConfigProvider: React.FC<{
  children: ReactNode;
}> = ({ children }) => {
  const [mapConfig, setMapConfig] = useState<IMapConfig>(defaultMapConfig);

  const handleChangeMapConfig = (mapConfig: Partial<IMapConfig>) => {
    setMapConfig((prev) => ({ ...prev, ...mapConfig }));
  };

  return (
    <MapConfigContext.Provider
      value={{ mapConfig, changeMapConfig: handleChangeMapConfig }}
    >
      {children}
    </MapConfigContext.Provider>
  );
};
