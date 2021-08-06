import React, { FC, useContext, useEffect, useState } from 'react';
import { GuestMyPageReservation } from '../../../../package/interface/Reservation';
import { getDay, getMonth, getHours } from '../../../../utils/function/GetDate';
import { UserContext } from '../../../../utils/context/UserContext';
import ReservationInfor from '../../parts/ReservationInfor/ReservationInfor';
import Title from '../parts/Title/Title';
import './Mypage.scss'
import { getGuestMypage } from '../../../../package/api';
import { Link, useHistory } from 'react-router-dom';

const Mypage: FC = () => {
  const [reserved, setReserved] = useState<GuestMyPageReservation>()
  const [previousReserved, setPreviousReserved] = useState<GuestMyPageReservation>()
  const { user } = useContext(UserContext)
  const history = useHistory()

  useEffect(() => {
    const getReserved = async () => {
      try {
        const response = await getGuestMypage()
        setReserved(response.reservations[0])
        setPreviousReserved(response.reservations[1])
      } catch (error) {
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
            {user.isBeautician ? <div className="mypage-profile-content">
              <Link to="/beautician/mypage">美容師用マイページ</Link>
            </div> : undefined}
          </div>
        </div>
        <div className="mypage-reserved-history">
          <ReservationInfor titleText="最新の予約" storeName={reserved ? reserved.salonName : ""} beauticianLastName={reserved ? reserved.beauticianLastName : ""}
            beauticianFirstName={reserved ? reserved.beauticianFirstName : ""} month={reserved ? getMonth(reserved.date) : null} day={reserved ? getDay(reserved.date) : null}
            hours={reserved ? getHours(reserved.date) : null} menus={reserved ? reserved.menus : []} />
          <ReservationInfor titleText="前回の予約" storeName={previousReserved ? previousReserved.salonName : ""} beauticianLastName={previousReserved ? previousReserved.beauticianLastName : ""}
            beauticianFirstName={previousReserved ? previousReserved.beauticianFirstName : ""} month={previousReserved ? getMonth(previousReserved.date) : null} day={previousReserved ? getDay(previousReserved.date) : null}
            hours={previousReserved ? getHours(previousReserved.date) : null} menus={previousReserved ? previousReserved.menus : []} />
        </div>
      </section>
    </div>
  )
}

export default Mypage;