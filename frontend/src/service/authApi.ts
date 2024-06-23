import { apiClient } from "./axiosApi";
import { LogInRequest, LogInResponse, SignUpRequest, SignUpResponse, ErrorResponse } from "@/types";
import { AuthError } from "@/errors/auth";
import { AxiosError } from "axios";
import { HTTP_CONFLICT, HTTP_UNAUTHORIZED } from "@/constants";

export const addNewUser = async (values: SignUpRequest): Promise<SignUpResponse> => {
    try {
        const response = await apiClient.post("/users", values);
        return response.data as SignUpResponse;
    } catch (error: unknown) {
        const err = error as AxiosError;
        console.error("Error making new user: ", err);
        if (err.response?.status === HTTP_CONFLICT) {
            throw new AuthError("User already exists", err.response?.status as number);
        }
        throw new AuthError("Failed to make new user", err.response?.status as number);
    }
}

export const logIn = async (values: LogInRequest): Promise<LogInResponse> => {
    try {
        const response = await apiClient.post("/auth/login", values);
        return response.data as LogInResponse;
    } catch (error: unknown) {
        const err: AxiosError = error as AxiosError;
        console.error("Error logging in: ", err);
        if (err.response?.status === HTTP_UNAUTHORIZED) {
            throw new AuthError("Invalid email or password", err.response?.status as number);
        }
        throw new Error("Failed to log in");
    }
}