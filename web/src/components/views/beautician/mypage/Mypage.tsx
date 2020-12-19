import React, { FC } from 'react';
import Title from '../../guest/parts/Title/Title';
import './Mypage.scss';

const Mypage: FC = () => {
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
              <h1>山本 浩一<span>ヤマモト コウイチ</span></h1>
              <button>
                変更
              </button>
            </span>
            <h2>メールアドレス</h2>
            <p>yamamoto@test.com</p>
            <h2>電話番号</h2>
            <p>09012344321</p>
            <div className="sns-content">
              <div className="line">
                <figure className="fab fa-line"></figure>
                <h2>yamamoto@line</h2>
              </div>
              <div className="instagram">
                <figure className="fab fa-instagram">
                </figure>
                <h2>yamamoto@line</h2>
              </div>
            </div>
          </div>
        </div>
        <div className="middle-content">
          <div className="menus">
            <h2>メニュー</h2>
            <dl>
              <span>
                <dt>カット</dt>
                <dd>3000</dd>
              </span>
              <span>
                <dt>カラー</dt>
                <dd>5000</dd>
              </span>
            </dl>
          </div>
          <div className="salons">
            <h2>美容院</h2>
            <dl>
              <span>
                <dt>カット</dt>
                <dd>3000</dd>
              </span>
              <span>
                <dt>カラー</dt>
                <dd>5000</dd>
              </span>
              <span>
                <dt>カット</dt>
                <dd>3000</dd>
              </span>
              <span>
                <dt>カラー</dt>
                <dd>5000</dd>
              </span>
            </dl>
          </div>
        </div>
      </div>
    </div>
  )
}

export default Mypage;