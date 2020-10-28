import React, { FC, useState } from 'react';
import Title from '../parts/Title/Title'
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
          return <ChooseCard key={i} type="beautician" image="/img/beautician_1.jpg" content={beautician} />
        })}
      </div>
    </div>
  )
}

export default ChooseBeautician;