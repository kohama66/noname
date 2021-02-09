import React, { FC, useEffect, useState } from 'react'
import { getSalonMypage } from '../../../../package/api'
import { salonMyPage } from '../../../../package/interface/response/Salon'
import Accordion from '../../parts/accordion/Accordion'
import Schedule from '../../parts/Schedule/Schedule'
import "./MyPage.scss"

const MyPage: FC = () => {
  const [mypage, setMypage] = useState<salonMyPage>()

  useEffect(() => {
    const handleGetMyPage = async () => {
      try {
        const response = await getSalonMypage("1")
        setMypage(response)
      } catch (error) {
        console.log(error)
      }
    }
    handleGetMyPage()
  }, [])

  return (
    <div id="mypage">
      <div className="contents">
        <figure>
          <img src="/img/salan.jpg" alt="" />
        </figure>
        <div className="profile">
          <div>
            <h1>{mypage?.salon.name}</h1>
            <span>
              <h2>PHONE</h2>
              <p>{mypage?.salon.phoneNumber}</p>
            </span>
            <span>
              <h2>POSTALCODE</h2>
              <p>{mypage?.salon.postalCode}</p>
            </span>
            <span>
              <h2>ADDRESS</h2>
              <p>{mypage ? `${mypage?.salon.prefectures}${mypage?.salon.city}${mypage?.salon.town}${mypage?.salon.addressCode}${mypage?.salon.addressOther}` : ""}</p>
            </span>
            <span>
              <h2>OPEN</h2>
              <p>{mypage?.salon.openingHours.slice(0, -3)}</p>
            </span>
            <span>
              <h2>CLOSE</h2>
              <p>{mypage?.salon.closingHours.slice(0, -3)}</p>
            </span>
          </div>
          <button>変更</button>
        </div>
        <div className="spaces">
          <h2>提供しているスペース</h2>
          <div>
            <i className="fas fa-chair"></i>
            <p>× 2</p>
          </div>
        </div>
        <div className="beauticians">
          <Accordion buttonText="STYLIST LIST▼">
            <ul>
              {mypage?.users.map(user => <li>{user.lastName + " " + user.firstName}</li>)}
            </ul>
          </Accordion>
        </div>

        <div className="reservation">
          <Schedule reservations={mypage ? mypage.reservations : []} />
        </div>
      </div>
    </div>
  )
}

export default MyPage
