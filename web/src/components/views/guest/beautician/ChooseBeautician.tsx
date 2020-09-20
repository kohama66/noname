import React, { FC, useState } from 'react';
import { BeauticianGetAll } from '../../../../package/interface/response/Reservation'
import BeauticianCard from './BeauticianCard'
import Title from '../parts/Title'

type props = {
  beauticians?: BeauticianGetAll
  // setState: (props: string) => void
}

const ChooseBeautician: FC<props> = ({ beauticians }) => {
  return (
    <div id="choose-beautician">
      <Title titleText={"美容師を選ぶ"} image={"/img/choose1.png"}/>
      <div className="contents">
        {beauticians?.beauticians.map((beautician, i) => {
          return <BeauticianCard lastName={beautician.lastName} key={i} randId={beautician.randId} />
        })}
      </div>
    </div>
  )
}

export default ChooseBeautician;