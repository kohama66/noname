import { MenuDetail } from "./Menu"

export interface Beautician {
	randId: string
	firstName: string
	lastName: string
	phoneNumber: string
	lineId: string
	instagramId: string
	comment: string
	menus: MenuDetail[]
}

export interface BeauticianGetAll {
  beauticians: Beautician[]
}

export const initBeautician: Beautician = <Beautician>{}

export const isBeauticianInterface = (arg: any): arg is Beautician => {
  return arg !== null &&
    typeof arg === "object" &&
    typeof arg.firstName === "string" && typeof arg.lastName === "string" && typeof arg.phoneNumber === "string" &&
    typeof arg.lineId === "string" && typeof arg.comment === "string" && typeof arg.instagramId === "string"
}