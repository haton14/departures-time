import { useState } from 'react';
import { StationApi } from '../api/api';
import DestinationComponent from './Destination';
import NearByStationComponent from './NearByStation';

const RouteComponent = (/*props: NearByStationProps*/) => {
  const [from, setFrom] = useState<string>('');
  const [to, setTo] = useState<string>('');
  const [routeURL, setRouteURL] = useState<string>('');

  const onChangeFrom = (e: any) => {
    setFrom(e.target.value);
    console.log(from);
  };

  const onChangeTo = (e: any) => {
    setTo(e.target.value);
    console.log(to);
  };

  const getRouteURL = async () => {
    if (from === '' || to === '') {
      alert('最寄駅と目的駅は必須');
      return;
    }
    const api = new StationApi();
    try {
      const url = await api.getV1Routes(from, to);
      setRouteURL(url.data.url);
    } catch (e: any) {
      alert(e.response.data.message);
    }
  };
  if (routeURL === '') {
    return (
      <>
        <NearByStationComponent onChange={onChangeFrom} from={from} />
        <br />
        <DestinationComponent onChange={onChangeTo} to={to} />
        <br />
        <button onClick={getRouteURL}>経路探索URL生成</button>
      </>
    );
  }
  return <a href={routeURL}>探索結果 for 駅すぱあとWebサービス</a>;
};

export default RouteComponent;
