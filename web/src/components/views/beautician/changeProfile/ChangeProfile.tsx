import React, { FC, useContext, useEffect, useState } from 'react';
import { UserContext } from '../../../../utils/context/UserContext';
import FormParts from '../../parts/formParts/formParts';
import './ChangeProfile.scss';

const ChangeProfile: FC = () => {
  const [firstName, setFirstName] = useState<string>("")
  const [lastName, setLastName] = useState<string>("")
  const [firstNameKana, setFirstNameKana] = useState<string>("")
  const [lastNameKana, setLastNameKana] = useState<string>("")
  const [phoneNumber, setPhoneNumber] = useState<string>("")
  const [disabled, setDisabled] = useState<boolean>(false)
  const [lineID, setLineID] = useState<string>("")
  const [instaID, setInstaID] = useState<string>("")

  const { user } = useContext(UserContext)
  useEffect(() => {
    setLastName(user.lastName)
    setFirstName(user.firstName)
    setLastNameKana(user.lastNameKana)
    setFirstNameKana(user.firstNameKana)
    setPhoneNumber(user.phoneNumber)
    if (user.beauticianInfo.lineId) {
      setLineID(user.beauticianInfo.lineId)
    }
    if (user.beauticianInfo.instagramId) {
      setInstaID(user.beauticianInfo.instagramId)
    }
  }, [])

  const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault()
    setDisabled(true)
  }

  return (
    <div id="change-profile">
      <form onSubmit={handleSubmit}>
        <FormParts label="苗字" type="text" value={lastName} setState={setLastName} required={true} disabled={disabled} />
        <FormParts label="名前" type="text" value={firstName} setState={setFirstName} required={true} disabled={disabled} />
        <FormParts label="苗字(フリガナ)" type="text" value={lastNameKana} setState={setLastNameKana} required={true} disabled={disabled} />
        <FormParts label="名前(フリガナ)" type="text" value={firstNameKana} setState={setFirstNameKana} required={true} disabled={disabled} />
        <FormParts label="電話番号" type="text" value={phoneNumber} setState={setPhoneNumber} minLength={11} maxLength={11} required={true} disabled={disabled} />
        <FormParts label="Line ID" type="text" value={lineID} setState={setLineID} required={true} disabled={disabled} />
        <FormParts label="Instagram ID" type="text" value={instaID} setState={setInstaID} required={true} disabled={disabled} />
        <FormParts type="submit" value="変更"  disabled={disabled} />
      </form>
    </div>
  )
}

export default ChangeProfile;