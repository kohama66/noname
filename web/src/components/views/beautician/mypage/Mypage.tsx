import React, { FC, useContext, useEffect, useState } from 'react';
import { useHistory } from 'react-router-dom';
import { getReservationBeautician, setHoliday } from '../../../../package/api';
import { Reservation } from '../../../../package/interface/Reservation';
import { UserContext } from '../../../../utils/context/UserContext';
import { dateToString } from '../../../../utils/function/GetDate';
import Title from '../../guest/parts/Title/Title';
import SingleForm from '../../parts/form/SingleForm';
import Modal from '../../parts/modal/Modal';
import Schedule from '../../parts/Schedule/Schedule';
import './Mypage.scss';

const Mypage: FC = () => {
  const [reserved, setReserved] = useState<Reservation[]>([])
  const [buttonDisabled, setButtonDisbled] = useState<boolean>(false)
  const [modalText, setModalText] = useState<string>("休日に設定しますか?")
  const [isModal, setIsModal] = useState<boolean>(false)
  const [clickDate, setClickDate] = useState<Date>()
  const { user } = useContext(UserContext)
  const history = useHistory()

  const handleModalClick = (date: Date) => {
    setIsModal(true)
    setClickDate(date)
  }

  const handleClick = async (type: "yes" | "no") => {
    setButtonDisbled(true)
    switch (type) {
      case "yes":
        try {
          if (clickDate) {
            await setHoliday({
              holiday: dateToString(clickDate)
            })
            handleGetReserved()
            setIsModal(false)
            setButtonDisbled(false)
          } else {
            setModalText("エラーが発生しました、しばらく後にやり直して下さい")
            setButtonDisbled(false)
          }
        } catch (error) {
          setModalText("エラーが発生しました、しばらく後にやり直して下さい")
          setButtonDisbled(false)
        }
        break
      case "no":
        setIsModal(false)
        setButtonDisbled(false)
        setModalText("休日に設定しますか?")
        break
    }
  }

  const handleGetReserved = async () => {
    const response = await getReservationBeautician()
    setReserved(response.reservations)
  }

  useEffect(() => {
    handleGetReserved()
  }, [])

  return (
    <div id="bt-mypage">
      <Title title="MY PAGE" text="マイページ" />
      <Modal modelText={modalText} isModal={isModal} handleClick={handleClick} buttonDisabled={buttonDisabled} />
      <div className="bt-mypage-contents">
        <div className="top-content">
          <figure>
            <img src="/img/beautician_1.jpg" alt="本人画像" />
          </figure>
          <div>
            <span>
              <h1>{user.lastName + " " + user.firstName}<span>{user.lastNameKana + " " + user.firstNameKana}</span></h1>
              <button onClick={() => history.push('/beautician/changeprofile')}>
                変更
              </button>
            </span>
            <h2>E-MAIL</h2>
            <p>{user.email}</p>
            <h2>PHONE-NUMBER</h2>
            <p>{user.phoneNumber}</p>
            <div className="sns-content">
              <div className="line">
                <figure className="fab fa-line"></figure>
                <h2>{user.beauticianInfo?.lineId ? user.beauticianInfo.lineId : "未設定"}</h2>
              </div>
              <div className="instagram">
                <figure className="fab fa-instagram">
                </figure>
                <h2>{user.beauticianInfo?.instagramId ? user.beauticianInfo.instagramId : "未設定"}</h2>
              </div>
            </div>
          </div>
        </div>
        <div className="middle-content">
          <div className="menus">
            <h2>MENUS</h2>
            <dl>
              {user.beauticianMenus?.map((menu, i) => {
                return <span key={i}>
                  <dt>{menu.name}</dt>
                  <dd>{menu.price}</dd>
                </span>
              })}
            </dl>
            <SingleForm disabled={buttonDisabled} type="menu" />
          </div>
          <div className="salons">
            <h2>SALONS</h2>
            <ul>
              {user.beauticianSalons?.map((salon, i) => {
                return <li key={i}>
                  <h3>{salon.name}</h3>
                  <p>{salon.prefectures + salon.city + salon.town + salon.addressOther}</p>
                </li>
              })}
            </ul>
          </div>
        </div>
        <div className="bottom-content">
          <Schedule reservations={reserved} handleModalClick={handleModalClick} />
        </div>
      </div>
    </div>
  )
}

export default Mypage;