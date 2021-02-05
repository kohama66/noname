import React, { FC, useContext, useEffect, useState } from 'react';
import { useHistory } from 'react-router-dom';
import { deleteBeauticianMenu, deleteBeauticianSalon, getBeauticianMypage, getMe, getReservationBeautician, setHoliday } from '../../../../package/api';
import { Reservation } from '../../../../package/interface/Reservation';
import { BeauticianMyPage } from '../../../../package/interface/response/Beautician';
import { dateToString } from '../../../../utils/function/GetDate';
import Title from '../../guest/parts/Title/Title';
import Accordion from '../../parts/accordion/Accordion';
import SingleForm from '../../parts/form/SingleForm';
import Modal from '../../parts/modal/Modal';
import Schedule from '../../parts/Schedule/Schedule';
import SearchSalon from '../../parts/searchSalon/SearchSalon';
import './Mypage.scss';

const Mypage: FC = () => {
  const [reserved, setReserved] = useState<Reservation[]>([])
  const [buttonDisabled, setButtonDisbled] = useState<boolean>(false)
  const [modalText, setModalText] = useState<string>("休日に設定しますか?")
  const [isModal, setIsModal] = useState<boolean>(false)
  const [clickDate, setClickDate] = useState<Date>()
  const [mypage, setMypage] = useState<BeauticianMyPage>()
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

  const handleDelete = async (randID: string, type: "menu" | "salon") => {
    try {
      switch (type) {
        case "menu":
          await deleteBeauticianMenu(randID)
          break
        case "salon":
          await deleteBeauticianSalon(randID)
          break
      }
      const response = await getBeauticianMypage()
      setMypage(response)
    } catch (error) {
      console.log(error.message)
    }
  }

  const handleGetReserved = async () => {
    try {
      const response = await getReservationBeautician()
      setReserved(response.reservations)
    } catch (error) {
      console.log(error)
    }
  }

  useEffect(() => {
    const hendleGetMypage = async () => {
      try {
        const response = await getBeauticianMypage()
        setMypage(response)
      } catch (error) {
        console.log(error)
      }
    }
    handleGetReserved()
    hendleGetMypage()
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
              <h1>{mypage?.user.lastName + " " + mypage?.user.firstName}<span>{mypage?.user.lastNameKana + " " + mypage?.user.firstNameKana}</span></h1>
              <button onClick={() => history.push('/beautician/changeprofile')}>
                変更
              </button>
            </span>
            <h2>EMAIL</h2>
            <p>{mypage?.user.email}</p>
            <h2>PHONE</h2>
            <p>{mypage?.user.phoneNumber}</p>
            <div className="sns-content">
              <div className="line">
                <figure className="fab fa-line"></figure>
                <h2>{mypage?.beautician?.lineId ? mypage?.beautician.lineId : "未設定"}</h2>
              </div>
              <div className="instagram">
                <figure className="fab fa-instagram">
                </figure>
                <h2>{mypage?.beautician.instagramId ? mypage?.beautician.instagramId : "未設定"}</h2>
              </div>
            </div>
          </div>
        </div>
        <div className="middle-content">
          <div className="menus">
            <Accordion buttonText="MENUS ▼">
              <dl>
                {mypage?.beauticianMenus.map((menu, i) => {
                  return <span key={i}>
                    <button className="delete-button" onClick={() => handleDelete(menu.randId, "menu")}>-</button>
                    <dt>{menu.name}</dt>
                    <dd>{menu.price}</dd>
                  </span>
                })}
              </dl>
              <SingleForm disabled={buttonDisabled} type="menu" />
            </Accordion>
          </div>
          <div className="salons">
            <Accordion buttonText="SALONS ▼" >
              <ul>
                {mypage?.salons.map((salon, i) => {
                  return <li key={i}>
                    <button className="delete-button" onClick={() => handleDelete(salon.randId, "salon")}>-</button>
                    <span>
                      <h3>{salon.name}</h3>
                      <p>{salon.prefectures + salon.city + salon.town + salon.addressOther}</p>
                    </span>
                  </li>
                })}
              </ul>
              <SearchSalon />
            </Accordion>
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