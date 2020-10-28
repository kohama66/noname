import React, { FC, useContext, useEffect, useState } from 'react';
import { findMenuDetails, getBeautician } from '../../../package/api';
import { Beautician, initBeautician } from '../../../package/interface/Beautician';
import { MenuDetail } from '../../../package/interface/Menu';
import { initSalon, Salon } from '../../../package/interface/Salon';
import { GeterSelectIDContext, GeterSelectValueContext } from '../guest';
import FinalComfirmationComponent from "./FinalComfirmation"

const FinalComfirmation: FC = () => {
  const [beautician, setBeautician] = useState<Beautician>(initBeautician)
  const [store, setStore] = useState<Salon>(initSalon)
  const [date, setDate] = useState<string>("")
  const [menus, setMenus] = useState<MenuDetail[]>([])

  const geterSelectValue = useContext(GeterSelectValueContext)
  const geterSelectID = useContext(GeterSelectIDContext)

  const handleGetAllSelected = async () => {
    const [beautician, store, _, date] = geterSelectValue()
    const menuIDs = geterSelectID("menu")
    try {
      if (typeof menuIDs === "object") {
        setMenus((await findMenuDetails(beautician.randId, menuIDs)).menuDetails)
        setBeautician(beautician)
        setStore(store)
        setDate(date)
      }
    } catch (error) {
      console.log(error)
    }
  }

  useEffect(() => {
    handleGetAllSelected()
  }, [])

  return (
    < FinalComfirmationComponent beautician={beautician} store={store} date={date} menus={menus} />
  )
}

export default FinalComfirmation;