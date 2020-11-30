import React, { FC, useContext, useEffect, useState } from 'react';
import { findSalons } from '../../../../package/api';
import { Salon } from '../../../../package/interface/Salon';
import { ReservedContext } from '../../../../utils/context/ReservadContext ';
import ChooseStoreComponent from './ChooseStore';

const ChooseStore: FC = () => {
  const [stores, setStores] = useState<Salon[]>([])

  const { getSelectID } = useContext(ReservedContext)

  const handleFindStores = async () => {
    const beauticianID = getSelectID("beautician")
    try {
      if (typeof beauticianID === "string" || beauticianID == null) {
        const response = await findSalons(beauticianID)
        setStores(response.salons)
      }
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