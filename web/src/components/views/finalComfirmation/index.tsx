import React, { FC, useContext, useEffect, useState } from 'react';
import { findMenuDetails, getBeautician } from '../../../package/api';
import { Beautician, initBeautician } from '../../../package/interface/Beautician';
import { MenuDetail } from '../../../package/interface/Menu';
import { initSalon, Salon } from '../../../package/interface/Salon';
import { GeterSelectIDContext } from '../guest';
import FinalComfirmationComponent from "./FinalComfirmation"

const FinalComfirmation: FC = () => {
  const [beautician, setBeautician] = useState<Beautician>(initBeautician)
  const [store, setStore] = useState<Salon>(initSalon)
  const [date, setDate] = useState<string>("")
  const [menus, setMenus] = useState<MenuDetail[]>([])

  const geterSelect = useContext(GeterSelectIDContext)

  const handleGetAllSelected = async () => {
    const beauticianID = geterSelect("beautician")
    // const storeID = geterSelect("store")
    const date = geterSelect("date")
    const menuIDs = geterSelect("menu")
    try {
      if (typeof beauticianID === "string" && typeof date === "string" && typeof menuIDs === "object") {
        setBeautician((await getBeautician(beauticianID)).beautician)
        setDate(date)
        setMenus((await findMenuDetails(beautician.randId, menuIDs)).menuDetails)
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