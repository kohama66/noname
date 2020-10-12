import React, { FC } from 'react';
// import { getMe, getReservationBeautician } from '../package/api';
import 'Form.scss'

const Form: FC = () => (
  <>
  <h2>登録する</h2>
  <form>
    <label>姓</label>
    <input type="text"/>
    <label>名</label>
    <input type="text"/>
    <label>メールアドレス</label>
    <input type="text"/>
  </form>
  {/* <button onClick={getReservationBeautician}>新規登録する</button> */}
  </>
);

export default Form;