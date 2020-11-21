import React, { FC, useState } from 'react';
import Input from '../parts/formParts/Input';
import './Login.scss'

const Login: FC = () => {
  const [email, setEmail] = useState<string>("")
  const [password, setPassword] = useState<string>("")
  const [disabled, setDisabled] = useState<boolean>(false)

  const handleSubmit = (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault()
    setDisabled(true)
  }

  return (
    <div id="login">
      <form onSubmit={handleSubmit}>
        <label>メールアドレス</label>
        <Input type="email" value={email} setState={setEmail} required={true} maxLength={50} disabled={disabled}/>
        <label>パスワード</label>
        <Input type="password" value={password} setState={setPassword} required={true} minLength={8}
          placeHolder="8文字以上で入力してください" disabled={disabled}/>
        <Input type="submit" value="ログイン" disabled={disabled}/>
      </form>
    </div>
  )
}

export default Login;