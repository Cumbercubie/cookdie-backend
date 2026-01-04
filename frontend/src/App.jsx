import "./App.css";
import 'leaflet/dist/leaflet.css';

import { MapContainer, TileLayer, Marker, Popup } from "react-leaflet";
import L from "leaflet";
import icon from "leaflet/dist/images/marker-icon.png";
import iconShadow from "leaflet/dist/images/marker-shadow.png";

// Fix for default marker icons
let DefaultIcon = L.icon({
  iconUrl: icon,
  shadowUrl: iconShadow,
  iconSize: [25, 41],
  iconAnchor: [12, 41],
});
L.Marker.prototype.options.icon = DefaultIcon;

function App() {
  const locations = [
    {
      id: "nyc",
      name: "NYC",
      coords: [40.7128, -74.006],
      options: ["Option 1", "Option 2"],
    },
    {
      id: "pa",
      name: "Pennsylvania",
      coords: [41.2033, -77.1945],
      options: ["Option A"],
    },
    {
      id: "nj",
      name: "New Jersey",
      coords: [40.0583, -74.4057],
      options: ["Option X", "Option Y"],
    },
  ];

  return (
    <>
      <MapContainer
        center={[39.96039145117331, -75.16866754775454]}
        style={{ height: "80vh", width: "90vw" }}
        zoom={14}
        scrollWheelZoom={true}
      >
        <TileLayer
          attribution='&copy; <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a> contributors'
          url="https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png"
        />
        <Marker position={[39.96039145117331, -75.16866754775454]}>
          <Popup>
            A pretty CSS3 popup. <br /> Easily customizable.
          </Popup>
        </Marker>
      </MapContainer>
    </>
  );
}

export default App;
