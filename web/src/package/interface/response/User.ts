import { User } from "../User";
import { GuestMyPageReservation } from "../Reservation"

export interface userResponse {
  user: User
}

export interface usersResponse {
  users: User[]
}

export interface guestMypageResponse {
  randId: string
  firstName: string
  lastName: string
  reservations: GuestMyPageReservation[]
}