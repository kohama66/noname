import React, { FC, useEffect, useState } from 'react';
import { getGuest } from '../../../../package/api';
import { GuestByMyPage, initGuest, initGuestByMyPage } from '../../../../package/interface/Guest';
import { GuestMyPageReservation } from '../../../../package/interface/Reservation';
import { getDay, getMonth, getHours } from '../../../../utils/GetDate';
import ReservationInfor from '../../parts/ReservationInfor/ReservationInfor';
import Title from '../parts/Title/Title';
import './Mypage.scss'

const Mypage: FC = () => {
  const [me, setMe] = useState<GuestByMyPage>(initGuestByMyPage)
  const [reserved, setReserved] = useState<GuestMyPageReservation>()
  const [previousReserved, setPreviousReserved] = useState<GuestMyPageReservation>()

  useEffect(() => {
    const handleGetMe = async () => {
      try {
        const response = await getGuest()
        setMe(response)
        if (response.reservations[0] != null) {
          setReserved(response.reservations[0])
        }
        if (response.reservations[1] != null) {
          setPreviousReserved(response.reservations[1])
        }
      } catch (error) {
        console.log(error)
      }
    }
    handleGetMe()
  }, [])

  return (
    <div id="guest-mypage">
      <Title title="MY PAGE" text="マイページ" />
      <section>
        <div className="profile-card">
          <figure></figure>
          <div>
            <dl>
              <span>
                <dt>名前</dt>
                <dd>{me.lastName + " " + me.firstName}</dd>
              </span>
              <span>
                <dt>歳</dt>
                <dd>29</dd>
              </span>
              <span>
                <dt>性別</dt>
                <dd>man</dd>
              </span>
              <span>
                <dt>電話</dt>
                <dd>09012345678</dd>
              </span>
            </dl>
          </div>
        </div>
        <div className="reserved-history-card">
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
                hours={(() => getHours(previousReserved.date))()} menus={previousReserved.menus}/>
            }
          })()
          }
        </div>
      </section>
    </div>
  )
}

export default Mypage;