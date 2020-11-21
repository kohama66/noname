import React, { FC, useState } from 'react';
import { Link } from 'react-router-dom';
import { signin } from '../../../package/api/auth';
import Title from '../guest/parts/Title/Title';
import Input from '../parts/formParts/Input';
import './Login.scss'

const Login: FC = () => {
  const [email, setEmail] = useState<string>("")
  const [password, setPassword] = useState<string>("")
  const [disabled, setDisabled] = useState<boolean>(false)

  const handleSubmit = async (e: React.FormEvent<HTMLFormElement>): Promise<void> => {
    e.preventDefault()
    setDisabled(true)
    await signin()
  }

  return (
    <div id="login">
      <Title title="LOGIN" text="ログイン" />
      <form onSubmit={handleSubmit}>
        <label>メールアドレス</label>
        <Input type="email" value={email} setState={setEmail} required={true} maxLength={50} disabled={disabled}/>
        <label>パスワード</label>
        <Input type="password" value={password} setState={setPassword} required={true} minLength={8}
          placeHolder="8文字以上で入力してください" disabled={disabled}/>
        <Input type="submit" value="ログイン" disabled={disabled}/>
        <Link to="signup" >未登録の方はこちら</Link>
      </form>
    </div>
  )
}

export default Login;