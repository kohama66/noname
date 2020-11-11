import React, { FC, useState } from 'react';
import Title from '../parts/Title/Title';
import './Mypage.scss'

const Mypage: FC = () => {
  const [me, setMe] = useState()
  
  return (
    <div id="guest-mypage">
      <Title title="MY PAGE" text="マイページ"/>
      <section>
        <div className="profile-card">
          <figure></figure>
          <div>
            <dl>
              <span>
                <dt>名前</dt>
                <dd>小濱 陵</dd>
              </span>
              <span>
                <dt>歳</dt>
                <dd>29</dd>
              </span>
              <span>
                <dt>性別</dt>
                <dd>man</dd>
              </span>
              <span>
                <dt>電話</dt>
                <dd>09012345678</dd>
              </span>
            </dl>
          </div>
        </div>
        <div className="reserved-history-card">
          <div>
            <h2>現在の予約</h2>
            <div>
              <dl>
                <span>
                  <dt>店舗</dt>
                  <dd>Kyoto Salon</dd>
                </span>
                <span>
                  <dt>美容師</dt>
                  <dd>山田 隆</dd>
                </span>
                <span>
                  <dt>日付</dt>
                  <dd>11月05日</dd>
                </span>
                <span>
                  <dt>時間</dt>
                  <dd>15:00から</dd>
                </span>
                <span className="menus">
                  <dt>メニュー</dt>
                  <span>
                    <dd>カット</dd>
                    <dd>カラー</dd>
                    <dd>パーマ</dd>
                  </span>
                </span>
                <span>
                  <dt>合計</dt>
                  <dd>5500</dd>
                </span>
              </dl>
            </div>
          </div>
          <div>
            <h2>前回の予約</h2>
            <div>
              <dl>
                <span>
                  <dt>店舗</dt>
                  <dd>Kyoto Salon</dd>
                </span>
                <span>
                  <dt>美容師</dt>
                  <dd>山田 隆</dd>
                </span>
                <span>
                  <dt>日付</dt>
                  <dd>11月05日</dd>
                </span>
                <span>
                  <dt>時間</dt>
                  <dd>15:00から</dd>
                </span>
                <span className="menus">
                  <dt>メニュー</dt>
                  <span>
                    <dd>カット</dd>
                    <dd>カラー</dd>
                    <dd>パーマ</dd>
                  </span>
                </span>
                <span>
                  <dt>合計</dt>
                  <dd>5500</dd>
                </span>
              </dl>
            </div>
          </div>
        </div>
      </section>
    </div>
  )
}

export default Mypage;