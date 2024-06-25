export interface GormModel {
    ID: number;
    CreatedAt: string;
    UpdatedAt: string;
    DeletedAt?: string;
}

export interface ReservationRequest {
    service: string;
    date: string;
    time: string;
}

export interface SignUpRequest {
    first_name: string;
    last_name?: string;
    email: string;
    password: string;
    phone_number: string;
}

export interface SignUpResponse extends GormModel {
    full_name: string;
    email: string;
    role: 'Customer' | 'Admin';
}

export interface LogInRequest {
    email: string;
    password: string;
}

export interface LogInResponse {
    token: string;
}

export interface ErrorResponse {
    error: string;
    error_code: number;
}

export interface Reservation {

}

export interface Review {
    full_name: string;
    rating: number;
    comment: string;
}

export interface User {
    full_name: string;
    email: string;
    phone_number: string;
    role: 'Customer' | 'Admin';
}