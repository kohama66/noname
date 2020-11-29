import React, { FC, memo } from 'react';
import { GuestContext, useGuestContext } from '../../../utils/context/GuestContext';

const Provider: FC = memo(({children}) => {
  const rootGuestContext = useGuestContext()

  return (
    <GuestContext.Provider value={rootGuestContext} >
      {children}
    </GuestContext.Provider>
  )
})

export default Provider;