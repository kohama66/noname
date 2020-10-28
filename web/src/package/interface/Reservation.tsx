export interface Reservation {
	date: string
	spaceId: number
	guestId: number
	beauticiaId: number
	menuId: number
}

export const isReservationInterface = (arg: any): arg is Reservation => {
	return arg !== null &&
		typeof arg === "object" &&
		typeof arg.date === "string" && typeof arg.spaceId === "number" && typeof arg.guestId === "number" &&
		typeof arg.beauticianId === "number" && typeof arg.menuId === "number"
}