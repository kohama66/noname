import React, { FC, useEffect, useState } from 'react';
import FormErrorMessage from '../formParts/FormErrorMessage';
import FormParts from '../formParts/formParts';
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
      <FormParts label="苗字" type="text" value={props.lastName} required={true} setState={props.setLastName} disabled={props.disabled} />
      <FormParts label="名前" type="text" value={props.firstName} required={true} setState={props.setFirstName} disabled={props.disabled} />
      <FormParts label="苗字(フリガナ)" type="text" value={props.lastNameKana} required={true} setState={props.setLastNameKana} disabled={props.disabled} />
      <FormParts label="名前(フリガナ)" type="text" value={props.firstNameKana} required={true} setState={props.setFirstNameKana} disabled={props.disabled} />
      <FormParts label="電話番号" type="text" value={props.phoneNumber} required={true} setState={props.setPhoneNumber} minLength={11} maxLength={11} disabled={props.disabled} />
      <FormParts label="メールアドレス" type="email" value={props.email} required={true} setState={props.setEmail} disabled={props.disabled} />
      <FormParts label="パスワード" type="password" value={props.password} required={true} setState={props.setPassword} minLength={8} placeHolder="8文字以上で入力してください" disabled={props.disabled} />
      <FormErrorMessage error={props.error} />
      <FormParts type="submit" value="登録" disabled={props.disabled} />
    </form>
  )
}

export default Form;