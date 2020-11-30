import React, { FC, useContext, useEffect, useState } from 'react';
import { findBeauticians } from '../../../../package/api';
import { Beautician } from '../../../../package/interface/Beautician';
import { ReservedContext } from '../../../../utils/context/ReservadContext ';
import ChooseBeauticianComponent from "./ChooseBeautician"

const ChooseBeautician: FC = () => {
  const [beauticians, setBeauticians] = useState<Beautician[]>([])
  const { getSelectID } = useContext(ReservedContext)

  const handleFindBeauticians = async () => {
    const storeID = getSelectID("store")
    const date = getSelectID("date")
    const menuIDs = getSelectID("menu")
    try {
      if ((typeof date === "string" || date == null) && (typeof storeID === "string" || storeID == null) && Array.isArray(menuIDs)) {
        const response = await findBeauticians(date, menuIDs, storeID)
        setBeauticians(response.beauticians)
      }
    } catch (err) {
      console.log(err)
    }
  }

  useEffect(() => {
    handleFindBeauticians()
  }, [])

  return (
    <ChooseBeauticianComponent beauticians={beauticians} />
  )
}

export default ChooseBeautician;