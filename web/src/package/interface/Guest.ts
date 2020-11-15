import { GuestMyPageReservation } from "./Reservation"

export interface Guest {
  randId: string
  firstName: string
  lastName: string
}

export interface GuestByMyPage {
  randId: string
  firstName: string
  lastName: string
  reservations: GuestMyPageReservation[]
}

export const initGuest: Guest = <Guest>{}
export const initGuestByMyPage: GuestByMyPage = <GuestByMyPage>{}
