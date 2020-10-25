import React, { FC, useContext, useEffect, useState } from 'react';
import { GeterSelectIDContext } from '..';
import { findBeauticians } from '../../../../package/api';
import { Beautician } from '../../../../package/interface/Beautician';
import ChooseBeauticianComponent from "./ChooseBeautician"

const ChooseBeautician: FC = () => {
  const [beauticians, setBeauticians] = useState<Beautician[]>([])
  const geterSelect = useContext(GeterSelectIDContext)

  const handleFindBeauticians = async () => {
    const storeID = geterSelect("store")
    const dateID = geterSelect("date")
    const menuIDs = geterSelect("menu")
    try {
      if (typeof dateID === "string" && typeof storeID === "string" && typeof menuIDs === "object") {
        const response = await findBeauticians(dateID, menuIDs, storeID)
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