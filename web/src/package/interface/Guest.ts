import { GuestMyPageReservation } from "./Reservation"

export interface Guest {
  randId: string
  firstName: string
  lastName: string
  firstNameKana: string
  lastNameKana: string
  email: string
  phoneNumber: string
}

export const initGuest: Guest = <Guest>{}