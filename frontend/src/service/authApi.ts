import { apiClient } from "./axiosApi";
import { LogInRequest, LogInResponse, SignUpRequest, SignUpResponse } from "@/types";

export const addNewUser = async (values: SignUpRequest): Promise<SignUpResponse> => {
    try {
        const response = await apiClient.post("/users", values);
        return response.data as SignUpResponse;
    } catch (error: unknown) {
        const err = error as Error;
        console.error("Error making new user: ", err.message);
        throw new Error("Failed to make new user");
    }
}

export const logIn = async (values: LogInRequest): Promise<LogInResponse> => {
    try {
        const response = await apiClient.post("/auth/login", values);
        return response.data as LogInResponse;
    } catch (error: unknown) {
        const err = error as Error;
        console.error("Error logging in: ", err.message);
        throw new Error("Failed to log in");
    }
}