export interface ReservationRequest {
    full_name: string;
    phoneNumber: string;
    service: string;
    date: string;
    time: string;
}

export interface Reservation {

}

export interface Review {
    full_name: string;
    rating: number;
    comment: string;
}

