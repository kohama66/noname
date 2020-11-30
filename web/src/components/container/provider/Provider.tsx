import React, { FC, memo } from 'react';
import { GuestContext, useGuestContext } from '../../../utils/context/GuestContext';
import { ReservedContext, useReservedContext } from '../../../utils/context/ReservadContext ';
// import useReservedContext, { ReservedContext } from '../../../utils/context/ReservadContext ';

const Provider: FC = memo(({ children }) => {
  const rootGuestContext = useGuestContext()
  const rootReservedContext = useReservedContext()
  return (
    <GuestContext.Provider value={rootGuestContext} >
      <ReservedContext.Provider value={rootReservedContext} >
        {children}
      </ReservedContext.Provider>
    </GuestContext.Provider>
  )
})

export default Provider;