import { MenuDetail } from "./Menu"
import { Salon } from "./Salon"
import { User } from "./User"
import { Menu } from "./Menu"

export interface Reservation {
	randId: string
	date: string
	holiday: boolean
	spaceId: number
	guestId: number
	beauticiaId: number
	menuId: number
}

export interface GuestMyPageReservation {
	id: number
	date: string
	salonName: string
	beauticianFirstName: string
	beauticianLastName: string
	menus: MenuDetail[]
}

export interface ReservationInfo {
	randID: string
	date: string
	salon: Salon
	user: User
	menus: Menu[]
}

export const isReservationInterface = (arg: any): arg is Reservation => {
	return arg !== null &&
		typeof arg === "object" &&
		typeof arg.date === "string" && typeof arg.spaceId === "number" && typeof arg.guestId === "number" &&
		typeof arg.beauticianId === "number" && typeof arg.menuId === "number"
}