import React, { FC, useContext, useEffect, useState } from 'react';
import { useHistory } from 'react-router-dom';
import { createReservation, findMenuDetails, getBeautician } from '../../../package/api';
import { initGuest } from '../../../package/interface/Guest';
import { MenuDetail } from '../../../package/interface/Menu';
import { GuestContext } from '../../../utils/context/GuestContext';
import { ReservedContext } from '../../../utils/context/ReservadContext ';
import { useError } from '../../../utils/hooks/Error';
import FinalComfirmationComponent from "./FinalComfirmation"

const FinalComfirmation: FC = () => {
  const { beautician, store, reservationDate, getMenuIDs } = useContext(ReservedContext)
  const [menus, setMenus] = useState<MenuDetail[]>([])
  const [totalPrice, setTotalPrice] = useState<number>(0)
  const history = useHistory()
  const { guest } = useContext(GuestContext)
  const { error, customError } = useError()

  const handleReserve = async () => {
    if (guest !== initGuest) {
      try {
        if (reservationDate != null) {
          await createReservation(beautician.randId, store.randId, getMenuIDs(), reservationDate)
          history.push("/reserved")
        }
      } catch (error) {
        console.log(error)
        customError("エラーです、内容をご確認下さい。")
      }
    } else {
      history.push("/login")
    }
  }

  useEffect(() => {
    const handleGetAllSelected = async () => {
      try {
        const menuDetails = await findMenuDetails(beautician.randId, getMenuIDs())
        setMenus(menuDetails.beauticianMenus)
        let totalPrice: number = 0
        menuDetails.beauticianMenus.map((detail) => {
          totalPrice += detail.price
        })
        setTotalPrice(totalPrice)
      } catch (error) {
        history.push("/guest")
      }
    }
    handleGetAllSelected()
  }, [])

  if (reservationDate != null) {
    return (
      < FinalComfirmationComponent beautician={beautician} store={store} date={reservationDate} menus={menus} totalPrice={totalPrice} handleReserve={handleReserve} error={error} />
    )
  } else {
    return <div></div>
  }
}

export default FinalComfirmation;