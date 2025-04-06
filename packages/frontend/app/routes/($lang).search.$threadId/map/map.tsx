import { useEffect, useRef } from "react";
import { MapContainer, Marker, Popup, TileLayer } from "react-leaflet";

import { useMapConfig } from "~/context/map-config-context";

import type { Map as IMap } from "leaflet";

import "leaflet/dist/leaflet.css";

// @see https://leaflet-extras.github.io/leaflet-providers/preview/ for the map layer
const Map = () => {
  const mapRef = useRef<IMap | null>(null);
  const {
    mapConfig: { center, zoom, markers },
  } = useMapConfig();

  useEffect(() => {
    if (mapRef.current) {
      mapRef.current.flyTo(center, zoom, { duration: 1 });
    }
  }, [center, zoom]);

  return (
    <MapContainer
      center={center}
      zoom={zoom}
      style={{ height: "100%", width: "100%" }}
      // Prevents infinite dragging across the globe, cuts the Pacific Ocean in half
      maxBounds={[
        [85, 180],
        [-85, -180],
      ]}
      ref={(el) => {
        mapRef.current = el;
      }}
    >
      <TileLayer
        url="https://server.arcgisonline.com/ArcGIS/rest/services/World_Street_Map/MapServer/tile/{z}/{y}/{x}"
        attribution='&copy; <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a> contributors'
      />
      {markers.map((marker) => (
        <Marker key={marker.name} position={marker.position}>
          <Popup>{marker.name}</Popup>
        </Marker>
      ))}
    </MapContainer>
  );
};

export default Map;
