import apiClient from "./axiosApi";
import { LogInRequest, LogInResponse, SignUpRequest, SignUpResponse, ErrorResponse } from "@/types";
import { AuthError } from "@/errors/auth";
import { AxiosError } from "axios";
import { HTTP_CONFLICT, HTTP_NOT_FOUND, HTTP_UNAUTHORIZED } from "@/constants";

export const logIn = async (values: LogInRequest): Promise<LogInResponse> => {
    try {
        const response = await apiClient.post("/auth/login", values);
        return response.data as LogInResponse;
    } catch (error: unknown) {
        const err: AxiosError = error as AxiosError;
        console.error("Error logging in: ", err);
        if (err.response?.status === HTTP_UNAUTHORIZED) {
            throw new AuthError("Invalid credentials", err.response?.status as number);
        }

        if (err.response?.status === HTTP_NOT_FOUND) {
            throw new AuthError("User not found", err.response?.status as number);
        }
        
        throw new Error("Failed to log in");
    }
}