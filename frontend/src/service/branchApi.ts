import apiClient from "./axiosApi";
import { type Branch } from "@/types";
import { getHourMinute } from "@/lib/utils";

export const addBranch = async (values: Branch) => {
  try {
    const response = await apiClient.post("/branches", values);
    return response.data.branch as Branch;
  } catch (error: unknown) {
    console.error("Error making new branch: ", error);
    throw new Error("Failed to make new branch");
  }
};

export const getBranches = async (): Promise<Branch[]> => {
  try {
    const response = await apiClient.get("/branches");
    return response.data.map((branch: Branch) => ({ 
      ...branch,
      open_time: getHourMinute(branch.open_time),
      close_time: getHourMinute(branch.close_time),
     })) as Branch[];
  } catch (error: unknown) {
    console.error("Error fetching branches: ", error);
    throw new Error("Failed to fetch branches");
  }
}

export const updateBranch = async (oldBranch: Branch, request: any): Promise<Branch> => {
  try {
    const response = await apiClient.put("/branches/" + oldBranch.ID, request);
    return response.data as Branch;
  } catch (error: unknown) {
    console.error("Error updating branch: ", error);
    throw new Error("Failed to update " + oldBranch.branch_name);
  }
}

export const deleteBranch = async (branch: Branch): Promise<void> => {
  try {
    await apiClient.delete("/branches/" + branch.ID);
  } catch (error: unknown) {
    console.error("Error deleting branch: ", error);
    throw new Error("Failed to delete " + branch.branch_name);
  }
}