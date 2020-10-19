import React, { FC, useContext, useEffect, useState } from 'react';
import { BeauticainIDContext } from '..';
import { findSalon } from '../../../../package/api';
import { Salon } from '../../../../package/interface/Salon';
import ChooseStoreComponent from './ChooseStore';

const ChooseStore: FC = () => {
  const [stores, setStores] = useState<Salon[]>([])

  const beauticianID = useContext(BeauticainIDContext)

  const findStores = async () => {
    try {
      const response = await findSalon(beauticianID)
      setStores(response.salons)
    } catch (err) {
      console.log(err)
    }
  }

  useEffect(() => {
    findStores()
  },[])

  return (
    <ChooseStoreComponent stores={stores} />
  )
}

export default ChooseStore;