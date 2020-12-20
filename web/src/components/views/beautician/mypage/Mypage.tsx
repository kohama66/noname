import React, { FC, useContext } from 'react';
import { UserContext } from '../../../../utils/context/UserContext';
import Title from '../../guest/parts/Title/Title';
import './Mypage.scss';

const Mypage: FC = () => {
  const { user } = useContext(UserContext)
  return (
    <div id="bt-mypage">
      <Title title="MY PAGE" text="マイページ" />
      <div className="bt-mypage-contents">
        <div className="top-content">
          <figure>
            <img src="/img/beautician_1.jpg" alt="本人画像" />
          </figure>
          <div>
            <span>
              <h1>{user.lastName + " " + user.firstName}<span>{user.lastNameKana + " " + user.firstNameKana}</span></h1>
              <button>
                変更
              </button>
            </span>
            <h2>メールアドレス</h2>
            <p>{user.email}</p>
            <h2>電話番号</h2>
            <p>{user.phoneNumber}</p>
            <div className="sns-content">
              <div className="line">
                <figure className="fab fa-line"></figure>
                <h2>{user.beauticianInfo.lineId ? user.beauticianInfo.lineId : "未設定"}</h2>
              </div>
              <div className="instagram">
                <figure className="fab fa-instagram">
                </figure>
                <h2>{user.beauticianInfo.instagramId ? user.beauticianInfo.instagramId : "未設定"}</h2>
              </div>
            </div>
          </div>
        </div>
        <div className="middle-content">
          <div className="menus">
            <h2>メニュー</h2>
            <dl>
              {user.beauticianMenus.map((menu, i) => {
                return <span key={i}>
                  <dt>{menu.name}</dt>
                  <dd>{menu.price}</dd>
                </span>
              })}
            </dl>
          </div>
          <div className="salons">
            <h2>美容院</h2>
            <ul>
              {user.beauticianSalons.map((salon, i) => {
                return <li key={i}>
                  <h3>{salon.name}</h3>
                  <p>{salon.prefectures + salon.city + salon.town + salon.addressOther}</p>
                </li>
              })}
            </ul>
          </div>
        </div>
      </div>
    </div>
  )
}

export default Mypage;