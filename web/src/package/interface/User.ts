import { GuestMyPageReservation } from "./Reservation"

export interface User {
  randId: string
  firstName: string
  lastName: string
  firstNameKana: string
  lastNameKana: string
  email: string
  phoneNumber: string
}

export const initUser: User = <User>{}