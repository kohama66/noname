import React, { FC, useContext, useEffect, useState } from 'react';
import Title from '../parts/Title/Title'
import './ChooseStore.scss'
import ChooseCard from '../parts/ChooseCard';
import { Salon } from '../../../../package/interface/Salon';
import { ReservedContext } from '../../../../utils/context/ReservadContext ';
import { findSalons } from '../../../../package/api';

const ChooseStore: FC = () => {
  const [stores, setStores] = useState<Salon[]>([])
  const { beautician } = useContext(ReservedContext)

  useEffect(() => {
    const handleFindStores = async () => {
      try {
        const response = await findSalons(beautician.randId)
        setStores(response.salons)
      } catch (err) {
        console.log(err)
      }
    }
    handleFindStores()
  }, [])
  
  return (
    <section id="choose-store">
      <Title title="SALON" text="お店から選ぶ" />
      <div className="choose-store-card-wrapper">
        {stores.map((store, id) => {
          return <ChooseCard image="/img/salan.jpg" type="store" content={store} key={id} />
        })}
      </div>
    </section>
  )
}

export default ChooseStore;