import React, { FC, useState } from 'react';
import { BeauticianGetAll } from '../../../../package/interface/response/Reservation'
import Title from '../parts/Title'
import "./ChooseBeautician.scss"
import ChooseCard from '../parts/ChooseCard/ChooseCard';

type props = {
  beauticians?: BeauticianGetAll
}

const ChooseBeautician: FC<props> = ({ beauticians }) => {
  return (
    <div id="choose-beautician">
      <Title title="BEAUTICIAN" text="美容師から選ぶ" />
      <div className="choose-beautician-wrapper">
        {/* <ChooseCard type="beautician" /> */}
      </div>
      {/* <div className="choose-beautician-contents">
        <h2 className="sub-title">BEAUTICIAN</h2>
        <div className="choose-beautician-contents-wrapper">
          {beauticians?.beauticians.map((beautician, i) => {
            return <BeauticianCard lastName={beautician.lastName} key={i} randId={beautician.randId} />
          })}
        </div>
      </div> */}
    </div>
  )
}

export default ChooseBeautician;