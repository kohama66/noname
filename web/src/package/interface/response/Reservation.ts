import { GuestMyPageReservation, Reservation, ReservationInfo } from "../Reservation";

export interface reservationsResponse {
  reservations: Reservation[]
}

export interface reservationResponse {
  reservation: Reservation
}

export interface guestMypageReservationsResponse {
  reservations: GuestMyPageReservation[]
}

export interface reservationInfoResponse {
  reservationInfo: ReservationInfo
}