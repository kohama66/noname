import { Guest } from "../Guest";
import { GuestMyPageReservation } from "../Reservation"

export interface guestResponse {
  user: Guest
}

export interface guestMypageResponse {
  randId: string
  firstName: string
  lastName: string
  reservations: GuestMyPageReservation[]
}