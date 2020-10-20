import React, { FC, useState } from 'react';
import Title from '../parts/Title'
import "./ChooseBeautician.scss"
import ChooseCard from '../parts/ChooseCard';
import { Beautician } from '../../../../package/interface/Beautician';

type props = {
  beauticians: Beautician[]
}

const ChooseBeautician: FC<props> = (props) => {
  return (
    <div id="choose-beautician">
      <Title title="BEAUTICIAN" text="美容師から選ぶ" />
      <div className="choose-beautician-wrapper">
        {props.beauticians.map((beautician, i) => {
          return <ChooseCard type="beautician" image="/img/beautician_1.jpg" content={beautician} />
        })}
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