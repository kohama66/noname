import React, { FC, useContext, useEffect, useState } from 'react';
import { useHistory } from 'react-router-dom';
import { getMe, updateBeautician } from '../../../../package/api';
import { UserContext } from '../../../../utils/context/UserContext';
import { useError } from '../../../../utils/hooks/Error';
import FormErrorMessage from '../../parts/formParts/FormErrorMessage';
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
  const history = useHistory()
  const { customError, error } = useError()
  const { user, setUser } = useContext(UserContext)

  
  useEffect(() => {
    setLastName(user.lastName)
    setFirstName(user.firstName)
    setLastNameKana(user.lastNameKana)
    setFirstNameKana(user.firstNameKana)
    setPhoneNumber(user.phoneNumber)
    if (user.beauticianInfo?.lineId) {
      setLineID(user.beauticianInfo.lineId)
    }
    if (user.beauticianInfo?.instagramId) {
      setInstaID(user.beauticianInfo.instagramId)
    }
  }, [])

  const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault()
    setDisabled(true)
    try {
      await updateBeautician({
        lastName: lastName,
        lastNameKana: lastNameKana,
        firstName: firstName,
        firstNameKana: firstNameKana,
        phoneNumber: phoneNumber,
        instagramId: instaID,
        lineId: lineID
      })
      const response = await getMe()
      setUser(response.user)
      history.push("/beautician/mypage")
    } catch (error) {
      customError(error.message)
      setDisabled(false)
    }
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
        <FormErrorMessage error={error} />
        <FormParts type="submit" value="変更" disabled={disabled} />
      </form>
    </div>
  )
}

export default ChangeProfile;