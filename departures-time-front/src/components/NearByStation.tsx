import { useState } from 'react';
import type { NearByStation } from '../api/api';
import { StationApi } from '../api/api';

interface NearByStationProps {
  onChange: (e: any) => void;
}

const NearByStationComponent = (props: NearByStationProps) => {
  const [nearByStations, setNearByStations] = useState<NearByStation[]>([]);

  const getCurrentPosition = () => {
    return new Promise((resolve: (value?: GeolocationPosition) => void, reject: (reason?: GeolocationPositionError) => void) => {
      navigator.geolocation.getCurrentPosition(resolve, reject);
    });
  };

  const getNearbyStations = async () => {
    const s = await getCurrentPosition();
    if (s !== undefined) {
      const api = new StationApi();
      try {
        const stations = await api.getV1NearbyStations(s.coords.longitude, s.coords.latitude);
        setNearByStations(stations.data.stations);
      } catch (e: any) {
        alert(e.response.data.message);
      }
    }
  };

  if (nearByStations.length === 0) {
    return (
      <button className="near-by-station" onClick={() => getNearbyStations()}>
        現在地取得
      </button>
    );
  }
  return (
    <select onChange={props.onChange}>
      {nearByStations.map((station) => (
        <option key={station.code} value={station.code}>
          {station.name} {station.distance}m
        </option>
      ))}
    </select>
  );
};

export default NearByStationComponent;
