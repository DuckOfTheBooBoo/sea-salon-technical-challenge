import { apiClient } from "./axiosApi";
import { SignUpRequest, SignUpResponse } from "@/types";

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