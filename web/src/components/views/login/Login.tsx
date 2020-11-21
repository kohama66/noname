import React, { FC } from 'react';
import Input from '../parts/formParts/Input';
import './Login.scss'

const Login: FC = () => {
  return (
    <div id="login">
      <form>
        <label>メールアドレス</label>
        <Input type="email" value=""/>
        <label>パスワード</label>
        <Input type="password" value="" />
        <Input type="submit" value="ログイン"/>
      </form>
    </div>
  )
}

export default Login;