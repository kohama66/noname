import React, { FC, memo } from 'react';
import { BeauticainContext, useBeauticianContext } from '../../../utils/context/BeauticianContext';
import { GuestContext, useGuestContext } from '../../../utils/context/GuestContext';
import { ReservedContext, useReservedContext } from '../../../utils/context/ReservadContext ';

const Provider: FC = memo(({ children }) => {
  const rootGuestContext = useGuestContext()
  const rootReservedContext = useReservedContext()
  const rootBeauticianContext = useBeauticianContext()
  return (
    <BeauticainContext.Provider value={rootBeauticianContext} >
      <GuestContext.Provider value={rootGuestContext} >
        <ReservedContext.Provider value={rootReservedContext} >
          {children}
        </ReservedContext.Provider>
      </GuestContext.Provider>
    </BeauticainContext.Provider>
  )
})

export default Provider;