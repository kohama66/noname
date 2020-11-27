import React, { createContext, FC, useState } from 'react';
import { Guest, initGuest } from '../../package/interface/Guest';

export const GuestContext = createContext({
  guest: initGuest,
  setGuest: (guest: Guest) => { },
})

export const useGuestContext = () => {
  const [guest, settingGuest] = useState<Guest>(initGuest)
  const setGuest = (guest: Guest): void => {
    settingGuest(guest)
  }
  return { guest, setGuest }
}
