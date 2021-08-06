export interface salonsRequest {
	beauticianRandId: string
}

export interface postBeauticianSalonRequest {
	salonRandId: string
}

export interface createSalonRequest {
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