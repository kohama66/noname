import React, { FC, useEffect, useState } from 'react';
import ChooseBeauticianComponet from '../../../../components/views/guest/beautician/ChooseBeautician'
import { getAllBeauticians } from '../../../../package/api/index'
import { BeauticianGetAll } from '../../../../package/interface/response/Reservation'

type props = {
  // setState: (props: string) => void 
}

const ChooseBeautician: FC<props> = (props) => {
  const [beauticians, setBeautician] = useState<BeauticianGetAll>()
  const handleBeautician = async () => {
    const response = await getAllBeauticians()
    setBeautician(response)
  }
  
  useEffect(() => {
    handleBeautician()
  }, [])
  return (
    <ChooseBeauticianComponet beauticians={beauticians} />
  )
}

export default ChooseBeautician;