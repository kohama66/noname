import React, { FC, useContext, useEffect, useState } from 'react';
import { GuestMyPageReservation } from '../../../../package/interface/Reservation';
import { getDay, getMonth, getHours } from '../../../../utils/function/GetDate';
import { UserContext } from '../../../../utils/context/UserContext';
import ReservationInfor from '../../parts/ReservationInfor/ReservationInfor';
import Title from '../parts/Title/Title';
import './Mypage.scss'
import { getGuestMypage } from '../../../../package/api';
import { initUser } from '../../../../package/interface/User';
import { useHistory } from 'react-router-dom';

const Mypage: FC = () => {
  const [reserved, setReserved] = useState<GuestMyPageReservation>()
  const [previousReserved, setPreviousReserved] = useState<GuestMyPageReservation>()
  const { user } = useContext(UserContext)
  const history = useHistory()

  useEffect(() => {
    const getReserved = async () => {
      if (user !== initUser) {
        try {
          const response = await getGuestMypage()
          setReserved(response.reservations[0])
          setPreviousReserved(response.reservations[1])
        } catch (error) {
          history.push("/login")
        }
      } else {
        history.push("/login")
      }
    }
    getReserved()
  }, [])

  return (
    <div id="guest-mypage">
      <Title title="MY PAGE" text="マイページ" />
      <section>
        <div className="mypage-profile">
          <h2>{`${user.lastName} ${user.firstName}`}</h2>
          <div className="mypage-profile-contents">
            <div className="mypage-profile-content">
              <h3>EMAIL</h3>
              <p>{user.email}</p>
            </div>
            <div className="mypage-profile-content">
              <h3>PHONE</h3>
              <p>{user.phoneNumber}</p>
            </div>
          </div>
        </div>
        <div className="mypage-reserved-history">
          {(() => {
            if (reserved) {
              return <ReservationInfor titleText="最新の予約" storeName={reserved.salonName} beauticianLastName={reserved.beauticianLastName}
                beauticianFirstName={reserved.beauticianFirstName} month={(() => getMonth(reserved.date))()} day={(() => getDay(reserved.date))()}
                hours={(() => getHours(reserved.date))()} menus={reserved.menus} />
            } else {
              return <ReservationInfor titleText="最新の予約" storeName="" beauticianFirstName="" beauticianLastName="" month={null} day={null} hours={null} menus={[]} />
            }
          })()
          }
          {(() => {
            if (previousReserved) {
              return <ReservationInfor titleText="前回の予約" storeName={previousReserved.salonName} beauticianLastName={previousReserved.beauticianLastName}
                beauticianFirstName={previousReserved.beauticianFirstName} month={(() => getMonth(previousReserved.date))()} day={(() => getDay(previousReserved.date))()}
                hours={(() => getHours(previousReserved.date))()} menus={previousReserved.menus} />
            } else {
              return <ReservationInfor titleText="前回の予約" storeName="" beauticianFirstName="" beauticianLastName="" month={null} day={null} hours={null} menus={[]} />
            }
          })()
          }
        </div>
      </section>
      {/* <section>
        <div>
          <div className="mypage-profile">
            <h2>プロフィール</h2>
            <dl>
              <span>
                <dd>名前</dd>
                <dt>{`${user.lastName} ${user.firstName} 様`}</dt>
              </span>
              <span>
                <dd>メール</dd>
                <dt>{user.email}</dt>
              </span>
              <span>
                <dd>電話</dd>
                <dt>{user.phoneNumber}</dt>
              </span>
            </dl>
          </div>
          <div className="beautician-mypage-button" >
            <h2>美容師用マイページ</h2>
          </div>
        </div>
        <div className="mypage-contents">
          {(() => {
            if (reserved) {
              return <ReservationInfor titleText="最新の予約" storeName={reserved.salonName} beauticianLastName={reserved.beauticianLastName}
                beauticianFirstName={reserved.beauticianFirstName} month={(() => getMonth(reserved.date))()} day={(() => getDay(reserved.date))()}
                hours={(() => getHours(reserved.date))()} menus={reserved.menus} />
            } else {
              return <ReservationInfor titleText="最新の予約" storeName="無し" beauticianFirstName="" beauticianLastName="" month={null} day={null} hours={null} menus={[]} />
            }
          })()
          }
          {(() => {
            if (previousReserved) {
              return <ReservationInfor titleText="前回の予約" storeName={previousReserved.salonName} beauticianLastName={previousReserved.beauticianLastName}
                beauticianFirstName={previousReserved.beauticianFirstName} month={(() => getMonth(previousReserved.date))()} day={(() => getDay(previousReserved.date))()}
                hours={(() => getHours(previousReserved.date))()} menus={previousReserved.menus} />
            } else {
              return <ReservationInfor titleText="最新の予約" storeName="無し" beauticianFirstName="" beauticianLastName="" month={null} day={null} hours={null} menus={[]} />
            }
          })()
          }
        </div>
      </section> */}
    </div>
  )
}

export default Mypage;