import React, { FC, useContext, useEffect, useState } from 'react';
import { findBeauticians } from '../../../../package/api';
import { Beautician } from '../../../../package/interface/Beautician';
import ChooseBeauticianComponent from "./ChooseBeautician"

const ChooseBeautician: FC = () => {
  const [beauticians, setBeauticians] = useState<Beautician[]>([])

  const handleFindBeauticians = async () => {
    try {
      const response = await findBeauticians()
      setBeauticians(response.beauticians)
    } catch (err) {
      console.log(err)
    }
  }
  
  useEffect(() => {
    handleFindBeauticians()
  },[])

  return (
    <ChooseBeauticianComponent beauticians={beauticians}/>
  )
}

export default ChooseBeautician;