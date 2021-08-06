import { Reservation } from "../Reservation";
import { Salon } from "../Salon";
import { User } from "../User";

export interface salonResponse {
  salon: Salon
}

export interface salonsResponse {
  salons: Salon[]
}

export interface salonMyPage {
  salon: Salon
  spaces: number
  reservations: Reservation[]
  users: User[]
}