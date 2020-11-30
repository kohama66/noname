import React, { FC, useContext, useEffect, useState } from 'react';
import { useHistory } from 'react-router-dom';
import { ReservedContext } from '../../../../utils/context/ReservadContext ';
import GuestComponent from './Guest'

const GuestHome: FC = () => {
  const { isAllChecked } = useContext(ReservedContext)
  const history = useHistory()
  useEffect(() => {
    if (isAllChecked()) {
      history.push("guest/final_comfirmation")
    }
  }, [])
  return (
    <GuestComponent />
  )
}

export default GuestHome;