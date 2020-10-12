import React, { FC, useState } from 'react';
import { BeauticianGetAll } from '../../../../package/interface/response/Reservation'
import BeauticianCard from './BeauticianCard'
import Title from '../parts/Title'
import "./ChooseBeautician.scss"

type props = {
  beauticians?: BeauticianGetAll
}

const ChooseBeautician: FC<props> = ({ beauticians }) => {
  return (
    <div id="choose-beautician">
      <Title titleText={"美容師を選ぶ"} image={"/img/choose1.png"} />
      <div className="choose-beautician-contents">
        <h2 className="sub-title">BEAUTICIAN</h2>
        <div className="choose-beautician-contents-wrapper">
          {beauticians?.beauticians.map((beautician, i) => {
            return <BeauticianCard lastName={beautician.lastName} key={i} randId={beautician.randId} />
          })}
        </div>
      </div>
    </div>
  )
}

export default ChooseBeautician;