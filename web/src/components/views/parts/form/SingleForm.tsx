import React, { FC, useState } from 'react';
import FormParts from '../formParts/formParts';
import Input from '../formParts/Input';
import Select from '../formParts/Select';
import "./SingleForm.scss";

interface props {
  type: "menu"
  disabled: boolean
}

const SingleForm: FC<props> = (props) => {
  const [stateName, setStateName] = useState<string>("")
  const [stateType, setStateType] = useState<string>("")

  return (
    <form id="single-form">
      <Input type="text" value={stateName} setState={setStateName} disabled={props.disabled} required={true} placeHolder="新しいメニューを追加する" />
      <Select type="menu" value={stateType} setState={setStateType} />
      <Input type="submit" value="登録" />
    </form>
  )

}

export default SingleForm;