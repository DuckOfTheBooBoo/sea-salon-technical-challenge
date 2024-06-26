import apiClient from "./axiosApi";
import { type Branch } from "@/types";

export const addBranch = async (values: Branch) => {
  try {
    const response = await apiClient.post("/branches", values);
    return response.data.branch as Branch;
  } catch (error: unknown) {
    console.error("Error making new branch: ", error);
    throw new Error("Failed to make new branch");
  }
};
