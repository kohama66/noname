import React, { FC } from 'react';
import Title from '../parts/Title/Title'
import './ChooseStore.scss'
import ChooseCard from '../parts/ChooseCard';
import { Salon } from '../../../../package/interface/Salon';

interface props {
  stores: Salon[]
}

const ChooseStore: FC<props> = (props) => {
  return (
    <section id="choose-store">
      <Title title="SALON" text="お店から選ぶ" />
      <div className="choose-store-card-wrapper">
        {props.stores.map((store, id) => {
          return <ChooseCard image="/img/salan.jpg" type="store" content={store} key={id} />
        })}
      </div>
    </section>
  )
}

export default ChooseStore;