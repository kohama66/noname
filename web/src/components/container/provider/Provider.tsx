import React, { FC, memo } from 'react';
import { BeauticainContext, useBeauticianContext } from '../../../utils/context/BeauticianContext';
import { UserContext, useUserContext } from '../../../utils/context/UserContext';
import { ReservedContext, useReservedContext } from '../../../utils/context/ReservadContext ';

const Provider: FC = memo(({ children }) => {
  const rootUserContext = useUserContext()
  const rootReservedContext = useReservedContext()
  const rootBeauticianContext = useBeauticianContext()
  return (
    <BeauticainContext.Provider value={rootBeauticianContext} >
      <UserContext.Provider value={rootUserContext} >
        <ReservedContext.Provider value={rootReservedContext} >
          {children}
        </ReservedContext.Provider>
      </UserContext.Provider>
    </BeauticainContext.Provider>
  )
})

export default Provider;