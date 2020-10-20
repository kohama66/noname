import React, { FC, useContext, useEffect, useState } from 'react';
import { BeauticainIDContext } from '..';
import { findSalons } from '../../../../package/api';
import { Salon } from '../../../../package/interface/Salon';
import ChooseStoreComponent from './ChooseStore';

const ChooseStore: FC = () => {
  const [stores, setStores] = useState<Salon[]>([])

  const beauticianID = useContext(BeauticainIDContext)

  const handleFindStores = async () => {
    try {
      const response = await findSalons(beauticianID)
      setStores(response.salons)
    } catch (err) {
      console.log(err)
    }
  }

  useEffect(() => {
    handleFindStores()
  },[])

  return (
    <ChooseStoreComponent stores={stores} />
  )
}

export default ChooseStore;