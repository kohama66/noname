import React, { FC, useState } from 'react';
import { BeauticianGetAll } from '../../../package/interface/response/Reservation'
import BeauticianCard from './BeauticianCard'

type props = {
  beauticians: BeauticianGetAll | undefined
}

const ChooseBeautician: FC<props> = ({ beauticians }) => {
  return (
    <div id="chooseBeautician">
      {beauticians?.Beauticians.map((beautician, i) => {
        return <BeauticianCard lastName={beautician.lastName} key={i} randId={beautician.randId}/>
      })}
    </div>
  )
}

export default ChooseBeautician;