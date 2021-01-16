import React, { FC, useContext, useEffect, useState } from 'react';
import "./SearchSalon.scss";
import { findSalonNoBelongs, findSalons, getMe, postBeauticianSalon } from '../../../../package/api';
import { Salon } from '../../../../package/interface/Salon';
import Input from '../formParts/Input';
import Select from '../formParts/Select';
import { UserContext } from '../../../../utils/context/UserContext';

const SearchSalon: FC = () => {
  const [salonRandID, setSalonRandID] = useState<string>("")
  const [salons, setSalons] = useState<Salon[]>([])
  const { setUser } = useContext(UserContext)

  const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault()
    try {
      await postBeauticianSalon({
        salonRandId: salonRandID
      })
      const response = await getMe()
      setUser(response.user)
      setSalonRandID("")
    } catch (error) {
      console.log(error)
    }
  }

  useEffect(() => {
    const handleFindSalons = async () => {
      try {
        const response = await findSalonNoBelongs()
        setSalons(response.salons)
      } catch (error) {
        console.log(error.message)
      }
    }
    handleFindSalons()
  }, [salonRandID])

  return (
    <form id="search-salon" onSubmit={handleSubmit}>
      <Select value={salonRandID} defaultValueLabel="新しい美容院の追加" optionData={salons} setState={setSalonRandID} required={true} />
      <Input type="submit" value="登録" />
    </form>
  )
}

export default SearchSalon;