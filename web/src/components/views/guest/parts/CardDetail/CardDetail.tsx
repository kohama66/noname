import React, { FC, useState } from 'react';
import { MenuDetail } from '../../../../../package/interface/Menu';

export interface StoreCardDetailProps {
  addressCode: string
  addressOther: string
  city: string
  closingHours: string
  name: string
  openingHours: string
  phoneNumber: string
  postalCode: string
  prefectures: string
  town: string
}

export const isStoreCardDetailProps = (arg: any): arg is StoreCardDetailProps => {
  return arg !== null &&
    typeof arg === "object" &&
    typeof arg.name === "string" && typeof arg.phoneNumber === "string" &&
    typeof arg.postalCode === "string" && typeof arg.openingHours === "string" && typeof arg.closingHours === "string" &&
    typeof arg.prefectures === "string" && typeof arg.city === "string" && typeof arg.town === "string" && typeof arg.addressCode && typeof arg.addressOther === "string"
}

export const StoreCardDetail: FC<StoreCardDetailProps> = (props) => {
  return (
    <div className="card-detail-store">
      <h3>{props.name}</h3>
      <p>{props.phoneNumber}</p>
      <p>{`営業時間: ${props.openingHours} ~ ${props.closingHours}`}</p>
      <p>{`〒 ${props.postalCode}`}</p>
      <p>{`${props.prefectures}${props.city}${props.town}${props.addressCode}${props.addressOther}`}</p>
    </div>
  )
}

export interface BeauticianCardDetailProps {
  firstName: string
  lasttName: string
  phoneNumber: string
  lineId?: string
  instagramId?: string
  menus: MenuDetail[]
}

export const BeauticianCardDetail: FC<BeauticianCardDetailProps> = (props) => {
  const [menuToggle, setMenuToggle] = useState<boolean>(false)

  const handleClickMenu = (e: React.MouseEvent<HTMLTableDataCellElement, MouseEvent>) => {
    e.stopPropagation()
    setMenuToggle(!menuToggle)
  }
  const handleClickInstagram = (e: React.MouseEvent<HTMLTableDataCellElement, MouseEvent>) => {
    e.stopPropagation()
  }

  return (
    <div>
      <h3>{props.lasttName + " " + props.firstName}</h3>
      <table>
        <tbody>
          <tr>
            <td>
              <p>LINE ID</p>
              <p>{props.lineId}</p>
            </td>
            <td>
              <p>Phone</p>
              <p>{props.phoneNumber}</p>
            </td>
          </tr>
          <tr>
            <td className="carddetail-menus" onClick={handleClickMenu} >
              <p>メニュー</p>
              <button>click</button>
              <ul className={"carddetail-menu" + (menuToggle ? " on-click" : "")}>
                {props.menus.map((menu, i) => {
                  return <li key={i}>
                    <p>{menu.name}</p>
                    <p>{menu.price}</p>
                  </li>
                })}
              </ul>
            </td>
            <td onClick={handleClickInstagram}>
              <p>Instagram</p>
              <object>
                <a href={props.instagramId} className="fab fa-instagram fa-2x" target="blank"></a>
              </object>
            </td>
          </tr>
        </tbody>
      </table>
    </div>
  )
}
