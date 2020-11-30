import React, { FC, useContext, useEffect, useState } from 'react';
import { useHistory } from 'react-router-dom';
import { createReservation, findMenuDetails, getBeautician } from '../../../package/api';
import { Beautician, initBeautician } from '../../../package/interface/Beautician';
import { initGuest } from '../../../package/interface/Guest';
import { MenuDetail } from '../../../package/interface/Menu';
import { initSalon, Salon } from '../../../package/interface/Salon';
import { GuestContext } from '../../../utils/context/GuestContext';
import { ReservedContext } from '../../../utils/context/ReservadContext ';
import { useError } from '../../../utils/hooks/Error';
import FinalComfirmationComponent from "./FinalComfirmation"

const FinalComfirmation: FC = () => {
  const { getSelectID, getSelectValue } = useContext(ReservedContext)
  const [beautician, setBeautician] = useState<Beautician>(initBeautician)
  const [store, setStore] = useState<Salon>(initSalon)
  const [date, setDate] = useState<string>("")
  const [menus, setMenus] = useState<MenuDetail[]>([])
  const [totalPrice, setTotalPrice] = useState<number>(0)
  const history = useHistory()
  const menuIDs = getSelectID("menu")
  const { guest } = useContext(GuestContext)
  const { error, customError } = useError()

  const handleGetAllSelected = async () => {
    const [beautician, store, _, date] = getSelectValue()
    try {
      if (typeof menuIDs === "object" && menuIDs.length !== 0 && beautician !== initBeautician) {
        const menuDetails = await findMenuDetails(beautician.randId, menuIDs)
        setMenus(menuDetails.beauticianMenus)
        setBeautician(beautician)
        setStore(store)
        setDate(date)
        let totalPrice: number = 0
        menuDetails.beauticianMenus.map((detail) => {
          totalPrice += detail.price
        })
        setTotalPrice(totalPrice)
      } else {
        history.push("/guest")
      }
    } catch (error) {
      console.log(error)
    }
  }
  const handleReserve = async () => {
    if (guest !== initGuest) {
      try {
        if (menuIDs != null && typeof menuIDs !== "string") {
          await createReservation(beautician.randId, store.randId, menuIDs, date)
          history.push("/reserved")
        }
      } catch (error) {
        console.log(error)
        customError("エラーです、内容をご確認下さい。")
      }
    } else {
      history.push("/guest/login")
    }
  }

  useEffect(() => {
    handleGetAllSelected()
  }, [])

  return (
    < FinalComfirmationComponent beautician={beautician} store={store} date={date} menus={menus} totalPrice={totalPrice} handleReserve={handleReserve} error={error} />
  )
}

export default FinalComfirmation;