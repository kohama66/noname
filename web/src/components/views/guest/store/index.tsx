import React, { FC, useContext, useEffect, useState } from 'react';
import { findSalons } from '../../../../package/api';
import { Salon } from '../../../../package/interface/Salon';
import { ReservedContext } from '../../../../utils/context/ReservadContext ';
import ChooseStoreComponent from './ChooseStore';

const ChooseStore: FC = () => {
  const [stores, setStores] = useState<Salon[]>([])
  const { beautician } = useContext(ReservedContext)

  const handleFindStores = async () => {
    try {
      const response = await findSalons(beautician.randId)
      setStores(response.salons)
    } catch (err) {
      console.log(err)
    }
  }

  useEffect(() => {
    handleFindStores()
  }, [])

  return (
    <ChooseStoreComponent stores={stores} />
  )
}

export default ChooseStore;