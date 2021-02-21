import React, { FC, useContext, useEffect, useState } from 'react';
import Title from '../parts/Title/Title'
import "./ChooseBeautician.scss"
import ChooseCard from '../parts/ChooseCard';
import { User } from '../../../../package/interface/User';
import { ReservedContext } from '../../../../utils/context/ReservadContext ';
import { findBeauticians } from '../../../../package/api';


const ChooseBeautician: FC = () => {
  const [beauticians, setBeauticians] = useState<User[]>([])
  const { store, reservationDate, getMenuIDs } = useContext(ReservedContext)

  useEffect(() => {
    const handleFindBeauticians = async () => {
      try {
        const response = await findBeauticians(reservationDate, getMenuIDs(), store.randId)
        setBeauticians(response.users)
      } catch (err) {
        console.log(err)
      }
    }
    handleFindBeauticians()
  }, [])

  return (
    <div id="choose-beautician">
      <Title title="BEAUTICIAN" text="美容師から選ぶ" />
      <div className="choose-beautician-wrapper">
        {beauticians.map((beautician, i) => {
          return <ChooseCard key={i} type="beautician" image="/img/beautician_1.jpg" content={beautician} />
        })}
      </div>
    </div>
  )
}

export default ChooseBeautician;