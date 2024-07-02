import { SignUpRequest, SignUpResponse, User } from "@/types";
import apiClient from "./axiosApi";
import { HTTP_CONFLICT } from "@/constants";
import { AuthError } from "@/errors/auth";
import { AxiosError } from "axios";

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

export const getUser = async (): Promise<User> => {
  try {
    const response = await apiClient.get("/reviews");
    return response.data.user as User;
  } catch (error: unknown) {
    console.error("Error fetching user information: ", error);
    throw new Error("Failed to fetch user information");
  }
}