import React, { FC, useEffect, useState } from 'react';
import ReservationInfor from '../parts/ReservationInfor/ReservationInfor';
import "./Reserved.scss";
import { getGuestMypage } from '../../../package/api';
import { GuestMyPageReservation } from '../../../package/interface/Reservation';
import { getDay, getHours, getMonth } from '../../../utils/function/GetDate';
import Title from '../guest/parts/Title/Title';

const Reserved: FC = () => {
  const [reserved, setReserved] = useState<GuestMyPageReservation>()

  useEffect(() => {
    const handleReserved = async () => {
      const response = await getGuestMypage()
      setReserved(response.reservations[0])
    }
    handleReserved()
  }, [])

  if (!reserved) return <div></div>
  return (
    <div id="reserved">
      <Title title="THANK YOU" text="予約完了" />
      <ReservationInfor titleText="今回の予約" storeName={reserved.salonName} beauticianLastName={reserved.beauticianLastName}
        beauticianFirstName={reserved.beauticianFirstName} month={(() => getMonth(reserved.date))()} day={(() => getDay(reserved.date))()}
        hours={(() => getHours(reserved.date))()} menus={reserved.menus} />
    </div>
  )
}

export default Reserved;