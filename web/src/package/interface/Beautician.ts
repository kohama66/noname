import { MenuDetail } from "./Menu"

export interface Beautician {
	randId: string
	firstName: string
	lastName: string
	firstNameKana: string
	lastNameKana: string
	phoneNumber: string
	beauticianInfo: BeauticianInfo
	menus?: MenuDetail[]
}

export interface BeauticianInfo {
	lineId?: string
	instagramId?: string
	comment?: string
}

export interface BeauticianGetAll {
  beauticians: Beautician[]
}

export const initBeautician: Beautician = <Beautician>{}

export const isBeauticianInterface = (arg: any): arg is Beautician => {
  return arg !== null &&
    typeof arg === "object" &&
    typeof arg.firstName === "string" && typeof arg.lastName === "string" && typeof arg.phoneNumber === "string"
}