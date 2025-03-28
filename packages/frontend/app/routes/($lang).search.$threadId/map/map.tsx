import { MapContainer, Marker, Popup, TileLayer } from "react-leaflet";

import "leaflet/dist/leaflet.css";

const Map = () => {
  return (
    <MapContainer
      center={[52.52437, 13.41053]}
      zoom={10}
      style={{ height: "100%", width: "100%" }}
    >
      <TileLayer
        url="https://tiles.stadiamaps.com/tiles/outdoors/{z}/{x}/{y}{r}.png"
        attribution='&copy; <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a> contributors'
      />
      <Marker position={[52.52437, 13.41053]}>
        <Popup>Berlin</Popup>
      </Marker>
    </MapContainer>
  );
};

export default Map;
