import { User } from "@/types";
import apiClient from "./axiosApi";

export const getUser = async (): Promise<User> => {
  try {
    const response = await apiClient.get("/reviews");
    return response.data.user as User;
  } catch (error: unknown) {
    console.error("Error fetching user information: ", error);
    throw new Error("Failed to fetch user information");
  }
}