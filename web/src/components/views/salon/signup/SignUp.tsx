import React, { FC, useState } from 'react'
import "./SignUp.scss"
import FormParts from '../../parts/formParts/formParts'
import Title from '../../guest/parts/Title/Title'
import { createSalon } from '../../../../package/api'

const SignUp: FC = () => {
  const [salonName, setSalonName] = useState<string>("")
  const [phoneNumner, setPhoneNumber] = useState<string>("")
  const [postalCode, setPostalCode] = useState<string>("")
  const [prefectures, setPrefectures] = useState<string>("")
  const [city, setCity] = useState<string>("")
  const [town, setTown] = useState<string>("")
  const [addressCode, setAddressCode] = useState<string>("")
  const [addressOther, setAddressOther] = useState<string>("")
  const [openHour, setOpenHour] = useState<string>("")
  const [closeHour, setCloseHour] = useState<string>("")

  const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault()
    try {
      const response = await createSalon({
        name: salonName,
        phoneNumber: phoneNumner,
        postalCode: postalCode,
        prefectures: prefectures,
        city: city,
        town: town,
        addressCode: addressCode,
        addressOther: addressOther,
        openingHours: openHour,
        closingHours: closeHour
      })
    } catch (error) {
      console.log(error)
    }
  }

  return (
    <div id="signup">
      <Title title="SIGNUP" text="美容院登録" />
      <form onSubmit={handleSubmit}>
        <FormParts type="text" value={salonName} setState={setSalonName} required={true} label="美容室名" />
        <FormParts type="text" value={postalCode} setState={setPostalCode} required={true} label="郵便番号" placeHolder="例)158-0001" maxLength={8} minLength={8} />
        <FormParts type="text" value={prefectures} setState={setPrefectures} required={true} label="都道府県" />
        <FormParts type="text" value={city} setState={setCity} required={true} label="市区" />
        <FormParts type="text" value={town} setState={setTown} required={true} label="町村" />
        <FormParts type="text" value={addressCode} setState={setAddressCode} required={true} label="番地" />
        <FormParts type="text" value={phoneNumner} setState={setPhoneNumber} required={true} label="電話番号" minLength={11} maxLength={11} />
        <FormParts type="text" value={addressOther} setState={setAddressOther} label="マンション名等" />
        <FormParts type="time" value={openHour} setState={setOpenHour} required={true} label="開店時刻" />
        <FormParts type="time" value={closeHour} setState={setCloseHour} required={true} label="閉店時刻" />
        <FormParts type="submit" value="登録" />
      </form>
    </div>
  )
}

export default SignUp
