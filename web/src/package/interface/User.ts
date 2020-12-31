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
  beauticianInfo?: BeauticianInfo
  beauticianMenus?: MenuDetail[]
  beauticianSalons?: Salon[]
  isBeautician: boolean
}

export const initUser: User = <User>{}

export const isUser = (arg: any): arg is User => {
  return arg !== null &&
    typeof arg === "object" &&
    typeof arg.firstName === "string" && typeof arg.lastName === "string" && typeof arg.phoneNumber === "string" &&
    typeof arg.firstNameKana === "string" && typeof arg.lastNameKana === "string" && typeof arg.email === "string" &&
    typeof arg.isBeautician === "boolean"
}