import React, { FC, useEffect, useState } from 'react';
import "./SearchSalon.scss";
import { findSalons } from '../../../../package/api';
import { Salon } from '../../../../package/interface/Salon';
import Input from '../formParts/Input';
import Select from '../formParts/Select';

const SearchSalon: FC = () => {
  const [salonRandID, setSalonRandID] = useState<string>("")
  const [salons, setSalons] = useState<Salon[]>([])

  const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault()

  }

  useEffect(() => {
    const handleFindSalons = async () => {
      try {
        const response = await findSalons()
        setSalons(response.salons)
      } catch (error) {
        console.log(error.message)
      }
    }
    handleFindSalons()
  }, [])

  return (
    <form id="search-salon" onSubmit={handleSubmit}>
      <Select value={salonRandID} defaultValueLabel="新しい美容院の追加" optionData={salons} setState={setSalonRandID} required={true} />
      <Input type="submit" value="登録" />
    </form>
  )
}

export default SearchSalon;