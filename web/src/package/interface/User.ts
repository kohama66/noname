import { BeauticianInfo } from "./Beautician"
import { MenuDetail } from "./Menu"
import { Salon } from "./Salon"

export interface User {
  randId: string
  firstName: string
  lastName: string
  firstNameKana: string
  lastNameKana: string
  email: string
  phoneNumber: string
  beauticianInfo: BeauticianInfo 
  beauticianMenus: MenuDetail[]
  beauticianSalons: Salon[]
  isBeautician: boolean
}

export const initUser: User = <User>{}