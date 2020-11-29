import React, { FC, useContext, useEffect, useState } from 'react';
import { getGuest } from '../../../package/api';
import { GuestContext } from '../../../utils/context/GuestContext';
import App from '../../App';

const Loading: FC = () => {
  const { guest, setGuest } = useContext(GuestContext)

  useEffect(() => {
    const checkGuest = async () => {
      try {
        const response = await getGuest()
        setGuest(response.guest)
      } catch (error) {
        console.log(error)
      }
    }
    checkGuest()
  }, [])

  return <App />
}

export default Loading;