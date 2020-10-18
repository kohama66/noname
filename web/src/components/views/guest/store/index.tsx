import React, { FC, useEffect, useState } from 'react';
import { findSalon } from '../../../../package/api';
import { salonsResponse } from '../../../../package/interface/response/Salon';
import { Salon } from '../../../../package/interface/Salon';
import ChooseStoreComponent from './ChooseStore';

const ChooseStore: FC = () => {
  const [stores, setStores] = useState<Salon[]>([])
  const handleSalon = async () => {
    try {
      const response = await findSalon()
      setStores(response.salons)
    } catch (err) {
      console.log(err)
    }
  }
  useEffect(() => {
    handleSalon()
  },[])
  return (
    <ChooseStoreComponent stores={stores} />
  )
}

export default ChooseStore;