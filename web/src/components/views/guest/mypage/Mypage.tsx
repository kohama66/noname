import React, { FC, useContext, useEffect, useState } from 'react';
import { GuestByMyPage, initGuest, initGuestByMyPage } from '../../../../package/interface/Guest';
import { GuestMyPageReservation } from '../../../../package/interface/Reservation';
import { getDay, getMonth, getHours } from '../../../../utils/function/GetDate';
import { GuestContext } from '../../../../utils/context/GuestContext';
import ReservationInfor from '../../parts/ReservationInfor/ReservationInfor';
import Title from '../parts/Title/Title';
import './Mypage.scss'

const Mypage: FC = () => {
  const [me, setMe] = useState<GuestByMyPage>(initGuestByMyPage)
  const [reserved, setReserved] = useState<GuestMyPageReservation>()
  const [previousReserved, setPreviousReserved] = useState<GuestMyPageReservation>()
  const { guest } = useContext(GuestContext)
 
  return (
    <div id="guest-mypage">
      <Title title="MY PAGE" text="マイページ" />
      <section>
        <div className="mypage-profile">
          <h2>{`${guest.lastName} ${guest.firstName} 様`}</h2>
          <dl>
            <span>
              <dd>メール</dd>
              <dt>guest_test@test.com</dt>
            </span>
            <span>
              <dd>電話</dd>
              <dt>09012344321</dt>
            </span>
          </dl>
        </div>
        <div className="mypage-contents">
          {(() => {
            if (reserved != null) {
              return <ReservationInfor titleText="今回の予約" storeName={reserved.salonName} beauticianLastName={reserved.beauticianLastName}
                beauticianFirstName={reserved.beauticianFirstName} month={(() => getMonth(reserved.date))()} day={(() => getDay(reserved.date))()}
                hours={(() => getHours(reserved.date))()} menus={reserved.menus} />
            }
          })()
          }
          {(() => {
            if (previousReserved != null) {
              return <ReservationInfor titleText="前回の予約" storeName={previousReserved.salonName} beauticianLastName={previousReserved.beauticianLastName}
                beauticianFirstName={previousReserved.beauticianFirstName} month={(() => getMonth(previousReserved.date))()} day={(() => getDay(previousReserved.date))()}
                hours={(() => getHours(previousReserved.date))()} menus={previousReserved.menus} />
            }
          })()
          }
        </div>
      </section>
    </div>
  )
}

export default Mypage;