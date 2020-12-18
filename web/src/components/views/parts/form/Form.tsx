import React, { FC } from 'react';
import Input from '../formParts/Input';
import './Form.scss'

interface props {
  handleSubmit: (e: React.FormEvent<HTMLFormElement>) => Promise<void>
  lastName: string
  firstName: string
  lastNameKana: string
  firstNameKana: string
  phoneNumber: string
  email: string
  password: string
  error: string | undefined
  disabled: boolean
  setLastName: React.Dispatch<React.SetStateAction<string>>
  setFirstName: React.Dispatch<React.SetStateAction<string>>
  setLastNameKana: React.Dispatch<React.SetStateAction<string>>
  setFirstNameKana: React.Dispatch<React.SetStateAction<string>>
  setPhoneNumber: React.Dispatch<React.SetStateAction<string>>
  setEmail: React.Dispatch<React.SetStateAction<string>>
  setPassword: React.Dispatch<React.SetStateAction<string>>
}

const Form: FC<props> = (props) => {
  return (
    <form onSubmit={props.handleSubmit}>
      <span>
        <label>苗字</label>
        <Input type="text" value={props.lastName} required={true} setState={props.setLastName} disabled={props.disabled} />
      </span>
      <span>
        <label>名前</label>
        <Input type="text" value={props.firstName} required={true} setState={props.setFirstName} disabled={props.disabled} />
      </span>
      <span>
        <label>苗字(フリガナ)</label>
        <Input type="text" value={props.lastNameKana} required={true} setState={props.setLastNameKana} disabled={props.disabled} />
      </span>
      <span>
        <label>名前(フリガナ)</label>
        <Input type="text" value={props.firstNameKana} required={true} setState={props.setFirstNameKana} disabled={props.disabled} />
      </span>
      <span>
        <label>電話番号</label>
        <Input type="text" value={props.phoneNumber} required={true} setState={props.setPhoneNumber} minLength={11} maxLength={11} disabled={props.disabled} />
      </span>
      <span>
        <label>メールアドレス</label>
        <Input type="email" value={props.email} required={true} setState={props.setEmail} disabled={props.disabled} />
      </span>
      <span>
        <label>パスワード</label>
        <Input type="password" value={props.password} required={true} setState={props.setPassword} minLength={8} placeHolder="8文字以上で入力してください" disabled={props.disabled} />
      </span>
      <span className="error-message">
        <p>{props.error}</p>
      </span>
      <span className="submit">
        <Input type="submit" value="登録" disabled={props.disabled} />
      </span>
    </form>
  )
}

export default Form;