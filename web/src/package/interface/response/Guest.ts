import { User } from "../User";
import { GuestMyPageReservation } from "../Reservation"

export interface guestResponse {
  user: User
}

export interface guestMypageResponse {
  randId: string
  firstName: string
  lastName: string
  reservations: GuestMyPageReservation[]
}