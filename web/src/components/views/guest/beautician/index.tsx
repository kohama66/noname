import React, { FC, useContext, useEffect, useState } from 'react';
import { findBeauticians } from '../../../../package/api';
import { User } from '../../../../package/interface/User';
import { ReservedContext } from '../../../../utils/context/ReservadContext ';
import ChooseBeauticianComponent from "./ChooseBeautician"

const ChooseBeautician: FC = () => {
  const [beauticians, setBeauticians] = useState<User[]>([])
  const { store, reservationDate, getMenuIDs } = useContext(ReservedContext)

  useEffect(() => {
    const handleFindBeauticians = async () => {
      try {
        const response = await findBeauticians(reservationDate, getMenuIDs(), store.randId)
        setBeauticians(response.users)
        console.log(response)
      } catch (err) {
        console.log(err)
      }
    }
    
    handleFindBeauticians()
  }, [])

  return (
    <ChooseBeauticianComponent beauticians={beauticians} />
  )
}

export default ChooseBeautician;