import React, { FC, useContext, useEffect, useState } from 'react';
import { useHistory } from 'react-router-dom';
import { getReservationInfo } from '../../../../../package/api';
import { Reservation } from '../../../../../package/interface/Reservation';
import { ReservedContext } from '../../../../../utils/context/ReservadContext ';
import SquareComponent from './Square'

interface props {
  day: Date
  time: number
  reservations: Reservation[]
  pathName?: string
}

const Square: FC<props> = (props) => {
  const history = useHistory()
  const [time, setTime] = useState<string>("")
  const [reserved, setReserved] = useState<Reservation>()
  const { setSelectValue } = useContext(ReservedContext)

  const handleClick = async () => {
    if (history.location.pathname === "/beautician/mypage") {
      if (reserved) {
        const response = await getReservationInfo(reserved.randId)
        const date = new Date(response.reservationInfo.date)
        console.log(date)
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
      } else {

      }
    } else {
      // ゲスト予約セット
      if (!reserved) {
        setTime(("0" + props.time).slice(-2) + ":00:00")
        setSelectValue(props.day.getFullYear() + "-" + ("0" + (props.day.getMonth() + 1)).slice(-2) + "-" + ("0" + props.day.getDate()).slice(-2) + " " + time)
        history.push("/guest")
      }
    }
  }

  useEffect(() => {
    const parseTime = (time: number) => {
      setTime(("0" + time).slice(-2) + ":00:00")
    }
    const verifyReserved = () => {
      props.reservations.map((reservation) => {
        if (reservation.date === props.day.getFullYear() + "-" + ("0" + (props.day.getMonth() + 1)).slice(-2) + "-" + ("0" + props.day.getDate()).slice(-2) + "T" + (("0" + props.time).slice(-2) + ":00:00Z")) {
          setReserved(reservation)
        }
      })
    }
    parseTime(props.time)
    verifyReserved()
  })

  return (
    <td onClick={handleClick} className={reserved && props.pathName !== "/beautician/mypage" ? "reserved-square" : ""}>
      {reserved ? "×" : ""}
    </td>
  )
}

export default Square;