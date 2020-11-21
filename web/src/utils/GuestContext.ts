import React, { createContext, FC, useState } from 'react';
import { Guest, initGuest } from '../package/interface/Guest';

export const GuestContext = createContext({
  guest: initGuest,
  settingGuest: (guest: Guest) => { },
})

export const useGuestContext = () => {
  const [guest, setGuest] = useState<Guest>(initGuest)
  const settingGuest = (guest: Guest): void => {
    setGuest(guest)
  }
  return { guest, settingGuest }
}
