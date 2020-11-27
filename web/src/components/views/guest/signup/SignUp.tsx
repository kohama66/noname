import React, { FC, FormEvent, useState } from 'react';
import { useHistory } from 'react-router-dom';
import { createGuest } from '../../../../package/api';
import { deleteToken, deleteUser, signup } from '../../../../package/api/auth';
import Input from '../../parts/formParts/Input';
import Title from '../parts/Title/Title';
import './SignUp.scss'

const SignUp: FC = () => {
  const [firstName, setFirstName] = useState<string>("")
  const [lastName, setLastName] = useState<string>("")
  const [firstNameKana, setFirstNameKana] = useState<string>("")
  const [lastNameKana, setLastNameKana] = useState<string>("")
  const [email, setEmail] = useState<string>("")
  const [passwoed, setPassword] = useState<string>("")
  const [disabled, setDisabled] = useState<boolean>(false)
  const history = useHistory()

  const handleSubmit = async (e: React.FormEvent<HTMLFormElement>): Promise<void> => {
    e.preventDefault()
    setDisabled(true)
    try {
      await signup({
        email: email,
        password: passwoed
      })
      try {
        const response = await createGuest({
          lastName: lastName,
          firstName: firstName,
          lastNameKana: lastNameKana,
          firstNameKana: firstNameKana,
          email: email
        })
      } catch (error) {
        deleteToken()
        deleteUser()
        throw new Error(error)
      }
      history.push("/guest")
    } catch (error) {
      console.log(error)
    }
    setDisabled(false)
  }

  return (
    <div id="signup">
      <Title title="SIGNUP" text="新規登録" />
      <form onSubmit={handleSubmit}>
        <span>
          <label>苗字</label>
          <Input type="text" value={lastName} required={true} setState={setLastName} disabled={disabled} />
        </span>
        <span>
          <label>名前</label>
          <Input type="text" value={firstName} required={true} setState={setFirstName} disabled={disabled} />
        </span>
        <span>
          <label>苗字(フリガナ)</label>
          <Input type="text" value={lastNameKana} required={true} setState={setLastNameKana} disabled={disabled} />
        </span>
        <span>
          <label>名前(フリガナ)</label>
          <Input type="text" value={firstNameKana} required={true} setState={setFirstNameKana} disabled={disabled} />
        </span>
        <span>
          <label>メールアドレス</label>
          <Input type="email" value={email} required={true} setState={setEmail} disabled={disabled} />
        </span>
        <span>
          <label>パスワード</label>
          <Input type="password" value={passwoed} required={true} setState={setPassword} minLength={8} placeHolder="8文字以上で入力してください" disabled={disabled} />
        </span>
        <span>
          <Input type="submit" value="登録" disabled={disabled} />
        </span>
      </form>
    </div>
  )
}

export default SignUp;