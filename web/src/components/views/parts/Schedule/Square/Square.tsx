import React, { FC, useContext, useEffect, useState } from 'react';
import './Square.scss';
import { useHistory } from 'react-router-dom';
import { getReservationInfo } from '../../../../../package/api';
import { Reservation } from '../../../../../package/interface/Reservation';
import { ReservedContext } from '../../../../../utils/context/ReservadContext ';
import { dateToString } from '../../../../../utils/function/GetDate';

interface props {
  day: Date
  time: number
  reservations: Reservation[]
  handleModalClick?: (date: Date) => void
}

const Square: FC<props> = (props) => {
  const history = useHistory()
  const [time, setTime] = useState<string>("")
  const [reserved, setReserved] = useState<Reservation>()
  const [holiday, setHoliday] = useState<boolean>(false)
  const { setSelectValue } = useContext(ReservedContext)


  const handleModalClick = () => {
    if (props.handleModalClick) {
      props.handleModalClick((() => {
        props.day.setHours(props.time)
        return props.day
      })())
    }
  }

  const handleClick = async () => {
    if (history.location.pathname === "/beautician/mypage") { // 美容師マイページでの動作
      if (reserved) { // 予約がある場合
        if (reserved.holiday) {
          handleModalClick()
        } else {
          const response = await getReservationInfo(reserved.randId)
          const date = new Date(response.reservationInfo.date)
          history.push({
            pathname: "/beautician/reservationverify",
            state: {
              userName: response.reservationInfo.user.lastName + " " + response.reservationInfo.user.firstName,
              userNameKana: response.reservationInfo.user.lastNameKana + " " + response.reservationInfo.user.firstNameKana,
              userPhoneNmber: response.reservationInfo.user.phoneNumber,
              salonName: response.reservationInfo.salon.name,
              salonAddress: response.reservationInfo.salon.prefectures + response.reservationInfo.salon.town +
                response.reservationInfo.salon.city + response.reservationInfo.salon.addressOther,
              reservedDate: `${date.getFullYear()}年 ${date.getMonth() + 1}月${date.getDate()}日`,
              reservedTime: `${date.getHours()}:${('0' + date.getMinutes()).slice(-2)}時から`,
              menus: response.reservationInfo.menus
            }
          })
        }
      } else { // 予約が無い場合
        handleModalClick()
      }
    } else {// ゲストページでの動作
      if (!reserved) {
        setTime(("0" + props.time).slice(-2) + ":00:00")
        setSelectValue(props.day.getFullYear() + "-" + ("0" + (props.day.getMonth() + 1)).slice(-2) + "-" + ("0" + props.day.getDate()).slice(-2) + " " + time)
        history.push("/guest")
      }
    }
  }
  useEffect(() => {
    const verifyReserved = () => {
      setReserved(undefined)
      props.day.setHours(props.time)
      props.reservations.map((reservation) => {
        if (reservation.date === dateToString(props.day)) {
          setReserved(reservation)
          if (reservation.holiday) {
            setHoliday(true)
          }
        }
      })
    }
    verifyReserved()

  }, [props.reservations])

  return (
    <td onClick={handleClick} className={(reserved ? "reserved-square " : " ") + (holiday ? "holiday-square" : "")}>
      {(() => {
        if (reserved) {
          if (holiday) {
            return "○"
          } else {
            return "×"
          }
        }
      })()}
    </td>
  )
}

export default Square;
