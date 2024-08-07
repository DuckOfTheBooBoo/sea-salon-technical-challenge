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
    branch_id: number;
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
    redirect_path: string;
}

export interface ErrorResponse {
    error: string;
    error_code: number;
}

export interface Reservation extends GormModel {
    full_name: string;
    phone_number: string;
    service: string;
    date: string;
    user_id: number;
    branch_id: number;
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

export interface Coordinate {
    lat: number;
    lng: number;
}

export interface Service {
    service_name: string;
    service_code: string;
    duration: number;
}

export interface Branch extends Coordinate, GormModel {
    branch_name: string;
    branch_address: string;
    services: Service[];
    open_time: string;
    close_time: string;
}

export interface JWTPayload {
    exp: number;
    iat: number;
    role: "Customer" | "Admin";
    ID: number;
    full_name: string;
}