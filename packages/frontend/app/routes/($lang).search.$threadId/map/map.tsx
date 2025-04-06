import { useEffect, useRef } from "react";
import { MapContainer, Marker, Popup, TileLayer } from "react-leaflet";

import type { Map as IMap, LatLngTuple } from "leaflet";

import "leaflet/dist/leaflet.css";

interface IMapProps {
  zoom?: number;
  center?: LatLngTuple;
}

// @see https://leaflet-extras.github.io/leaflet-providers/preview/ for the map layer
const Map = ({ center = [52.52437, 13.41053], zoom = 10 }: IMapProps) => {
  const mapRef = useRef<IMap | null>(null);

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
      ref={(el) => {
        mapRef.current = el;
      }}
    >
      <TileLayer
        url="https://tiles.stadiamaps.com/tiles/outdoors/{z}/{x}/{y}{r}.png"
        attribution='&copy; <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a> contributors'
      />
      <Marker position={center}>
        <Popup>Berlin</Popup>
      </Marker>
    </MapContainer>
  );
};

export default Map;
