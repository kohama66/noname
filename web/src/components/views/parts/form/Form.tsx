import React, { FC } from 'react';
import './Form.scss'

interface props {
  handleSubmit: (e: React.FormEvent<HTMLFormElement>) => Promise<void>
}

const Form: FC<props> = (props) => {
  return (
    <form onSubmit={props.handleSubmit}>
      {/* <label>メールアドレス</label>
      <Input type="email" value={email} setState={setEmail} required={true} maxLength={50} disabled={disabled} />
      <label>パスワード</label>
      <Input type="password" value={password} setState={setPassword} required={true} minLength={8}
        placeHolder="8文字以上で入力してください" disabled={disabled} />
      <Input type="submit" value="ログイン" disabled={disabled} />
      <Link to="signup" >未登録の方はこちら</Link> */}
    </form>
  )
}

export default Form;