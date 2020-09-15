import React, { FC } from 'react';
import { BeauticianGetAll } from '../../../package/interface/response/Reservation'

type props = {
  beauticians: BeauticianGetAll | undefined
}

const ChooseBeautician: FC<props> = ({ beauticians }) => {
  return (
    <div>
      {beauticians?.Beauticians.map((beautician) => {
        return <p>{beautician.lastName}</p>
      })}
    </div>
  )
}

export default ChooseBeautician;