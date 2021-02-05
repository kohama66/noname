import React, { FC, memo } from 'react';
import { UserContext, useUserContext } from '../../../utils/context/UserContext';
import { ReservedContext, useReservedContext } from '../../../utils/context/ReservadContext ';

const Provider: FC = memo(({ children }) => {
  const rootUserContext = useUserContext()
  const rootReservedContext = useReservedContext()
  return (
    <UserContext.Provider value={rootUserContext} >
      <ReservedContext.Provider value={rootReservedContext} >
        {children}
      </ReservedContext.Provider>
    </UserContext.Provider>
  )
})

export default Provider;