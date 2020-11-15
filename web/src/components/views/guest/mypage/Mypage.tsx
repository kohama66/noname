import React, { FC, useEffect, useState } from 'react';
import { getGuest } from '../../../../package/api';
import { GuestByMyPage, initGuestByMyPage } from '../../../../package/interface/Guest';
import ReservationInfor from '../../parts/ReservationInfor/ReservationInfor';
import Title from '../parts/Title/Title';
import './Mypage.scss'

const Mypage: FC = () => {
  const [me, setMe] = useState<GuestByMyPage>(initGuestByMyPage)
  const handleGetMe = async () => {
    try {
      const response = await getGuest()
      setMe(response)
      console.log(response)
    } catch (error) {
      console.log(error)
    }
  }
  useEffect(() => {
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
            if (me != initGuestByMyPage) {
              return <ReservationInfor titleText="今回の予約" storeName={me.reservations[0].salonName}
                beauticianLastName={me.reservations[0].beauticianLastName} beauticianFirstName={me.reservations[0].beauticianFirstName}
                month={(() => {
                  const month = new Date(me.reservations[0].date)
                  return month.getMonth()
                })()}
                day={(() => {
                  const day = new Date(me.reservations[0].date)
                  return day.getDay()
                })()}
                time={(() => {
                  const time = new Date(me.reservations[0].date)
                  return time.getHours()
                })()} />
            }
          })()}
          {(() => {
            if (me != initGuestByMyPage) {
              return <ReservationInfor titleText="前回の予約" storeName={me.reservations[1].salonName}
                beauticianLastName={me.reservations[1].beauticianLastName} beauticianFirstName={me.reservations[1].beauticianFirstName}
                month={(() => {
                  const month = new Date(me.reservations[1].date)
                  console.log(month)
                  return month.getMonth()
                })()}
                day={(() => {
                  const day = new Date(me.reservations[1].date)
                  return day.getDay()
                })()}
                time={(() => {
                  const time = new Date(me.reservations[1].date)
                  return time.getHours()
                })()} />
            }
          })()}
        </div>
      </section>
    </div>
  )
}

export default Mypage;