import React, { FC, useState } from 'react'
import "./SignUp.scss"
import FormParts from '../../parts/formParts/formParts'
import Title from '../../guest/parts/Title/Title'

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

  return (
    <div id="signup">
      <Title title="SIGNUP" text="美容院登録" />
      <form>
        <FormParts type="text" value={salonName} setState={setSalonName} required={true} label="美容室名" />
        <FormParts type="text" value={postalCode} setState={setPostalCode} required={true} label="郵便番号" placeHolder="ハイフン無しで入力して下さい" maxLength={7} minLength={7} />
        <FormParts type="text" value={prefectures} setState={setPrefectures} required={true} label="都道府県" />
        <FormParts type="text" value={city} setState={setCity} required={true} label="市区" />
        <FormParts type="text" value={town} setState={setTown} required={true} label="町村" />
        <FormParts type="text" value={addressCode} setState={setAddressCode} required={true} label="番地" />
        <FormParts type="text" value={phoneNumner} setState={setPhoneNumber} required={true} label="電話番号" />
        <FormParts type="text" value={addressOther} setState={setAddressOther} label="マンション名等" />
        <FormParts type="time" value={openHour} setState={setOpenHour} required={true} label="開店時刻" />
        <FormParts type="time" value={closeHour} setState={setCloseHour} required={true} label="閉店時刻" />
        <FormParts type="submit" value="登録" />
      </form>
    </div>
  )
}

export default SignUp
