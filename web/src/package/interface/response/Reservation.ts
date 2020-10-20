export interface Reservation {
  date: string;
  time: string;
}

export interface ReservationFindByBeautician {
  Reservations: Reservation[]
}
