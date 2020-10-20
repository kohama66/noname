export interface Beautician {
	randId: string
	firstName: string
	lastName: string
	phoneNumber: string
	lineId: string
	instagramId: string
	comment: string
}

export interface BeauticianGetAll {
  beauticians: Beautician[]
}
