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
    if (from === '' && to === '') {
      // エラー処理は省略
      return;
    }
    const api = new StationApi();
    const url = await api.getV1Routes(from, to);
    if (url.status !== 200) {
      // エラー処理は省略
      return;
    }
    setRouteURL(url.data.url);
  };
  if (routeURL === '') {
    return (
      <>
        <NearByStationComponent onChange={onChangeFrom} />
        <DestinationComponent onChange={onChangeTo} />
        <button onClick={getRouteURL}>経路探索URL生成</button>
      </>
    );
  }
  return <a href={routeURL}>探索結果 for 駅すぱあとWebサービス</a>;
};

export default RouteComponent;
