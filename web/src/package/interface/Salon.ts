export interface Salon {
	randId: string
	name: string
	phoneNumber: string
	openingHours: string
	closingHours: string
	postalCode: string
	prefectures: string
	city: string
	town: string
	addressCode: string
	addressOther: string
}

export const initSalon: Salon = <Salon>{}

export const isSalonInterface = (arg: any): arg is Salon => {
  return arg !== null &&
    typeof arg === "object" &&
		typeof arg.randId === "string" && typeof arg.name === "string" && typeof arg.phoneNumber === "string" &&
		typeof arg.openingHours === "string" && typeof arg.closingHours === "string" && typeof arg.postalCode === "string" &&
		typeof arg.prefectures === "string" && typeof arg.city === "string" && typeof arg.town === "string" && typeof arg.addressCode === "string" &&
		typeof arg.addressOther === "string"
}