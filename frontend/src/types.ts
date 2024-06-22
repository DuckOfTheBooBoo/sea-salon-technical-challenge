export interface GormModel {
    ID: number;
    CreatedAt: string;
    UpdatedAt: string;
    DeletedAt?: string;
}

export interface ReservationRequest {
    full_name: string;
    phoneNumber: string;
    service: string;
    date: string;
    time: string;
}

export interface SignUpRequest {
    first_name: string;
    last_name?: string;
    email: string;
    password: string;
    phoneNumber: string;
}

export interface SignUpResponse extends GormModel {
    full_name: string;
    email: string;
    role: 'Customer' | 'Admin';
}

export interface Reservation {

}

export interface Review {
    full_name: string;
    rating: number;
    comment: string;
}

