import { User } from "./User"

export interface BeauticianInfo {
	lineId?: string
	instagramId?: string
	comment?: string
}

// export const isBeautician = (arg: any): arg is User => {
//   return arg !== null &&
//     typeof arg === "object" &&
//     typeof arg.firstName === "string" && typeof arg.lastName === "string" && typeof arg.phoneNumber === "string" &&
//     typeof arg.firstNameKana === "string" && typeof arg.lastNameKana === "string" && typeof arg.email === "string" &&
// 		typeof arg.isBeautician === "boolean" && typeof arg.beauticianInfo === "object" &&
// 		typeof arg.beauticianMenus === "object" && Array.isArray(arg.beauticianMenus) && 
// 		typeof arg.beauticianSalons === "object" && Array.isArray(arg.beauticianSalons)
// }

export const isBeautician = (arg: any): arg is User => {
	return arg !== null &&
		typeof arg === "object" &&
		typeof arg.firstName === "string" && typeof arg.lastName === "string" && typeof arg.phoneNumber === "string" &&
		typeof arg.firstNameKana === "string" && typeof arg.lastNameKana === "string" && typeof arg.email === "string" &&
		typeof arg.isBeautician === "boolean" && typeof arg.beauticianInfo === "object" &&
		typeof arg.beauticianMenus === "object"
}