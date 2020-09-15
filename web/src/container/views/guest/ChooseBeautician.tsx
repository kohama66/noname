import React, { FC, useEffect, useState } from 'react';
import ChooseBeauticianComponet from '../../../components/views/guest/ChooseBeautician'
import { getAllBeauticians } from '../../../package/api/index'
import { BeauticianGetAll } from '../../../package/interface/response/Reservation'

const ChooseBeautician: FC = () => {
  const [beauticians, setBeautician] = useState<BeauticianGetAll>()
  const handleBeautician = async () => {
    const response = await getAllBeauticians()
    setBeautician(response)
  }
  useEffect(() => {
    handleBeautician()
  },[])
  return(
    <ChooseBeauticianComponet beauticians={beauticians} />
  )
}

export default ChooseBeautician;