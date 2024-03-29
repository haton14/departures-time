import { useState } from 'react';
import type { Destination } from '../api/api';
import { StationApi } from '../api/api';

interface DestinationProps {
  to: string;
  onChange: (e: any) => void;
}

const DestinationComponent = (props: DestinationProps) => {
  const [destinations, setDestinations] = useState<Destination[]>([]);
  const [name, setName] = useState<string>('');

  const getDestinations = async () => {
    const api = new StationApi();
    try {
      const stations = await api.getV1Destinations(name);
      setDestinations(stations.data.stations);
    } catch (e: any) {
      alert(e.response.data.message);
    }
  };

  if (destinations.length === 0) {
    return (
      <div>
        <input
          type="text"
          value={name}
          onChange={(e: any) => {
            setName(e.currentTarget.value);
          }}
        />
        <button onClick={getDestinations}>目的地駅検索</button>
      </div>
    );
  }
  return (
    <select onChange={props.onChange} value={props.to}>
      <option value="">目的駅を選択してください</option>
      {destinations.map((station) => (
        <option key={station.code} value={station.code}>
          {station.name}
        </option>
      ))}
    </select>
  );
};

export default DestinationComponent;
