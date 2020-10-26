import React, { FC, useContext, useEffect, useState } from 'react';
import { getBeautician } from '../../../package/api';
import { Beautician, initBeautician } from '../../../package/interface/Beautician';
import { initSalon, Salon } from '../../../package/interface/Salon';
import { GeterSelectIDContext } from '../guest';
import FinalComfirmationComponent from "./FinalComfirmation"

const FinalComfirmation: FC = () => {
  const [beautician, setBeautician] = useState<Beautician>(initBeautician)
  const [store, setStore] = useState<Salon>(initSalon)
  const [date, setDate] = useState<string>("")
  const geterSelect = useContext(GeterSelectIDContext)

  const handleGetAllSelected = async () => {
    const beauticianID = geterSelect("beautician")
    const storeID = geterSelect("store")
    const date = geterSelect("date")
    try {
      if (typeof beauticianID === "string" && typeof date === "string") {
        setBeautician((await getBeautician(beauticianID)).beautician)
        setDate(date)
      }
    } catch (error) {
      console.log(error)
    }
  }

  useEffect(() => {
    handleGetAllSelected()
  },[])

  return (
    < FinalComfirmationComponent beautician={beautician} store={store} date={date}/>
  )
}

export default FinalComfirmation;