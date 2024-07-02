import apiClient from "./axiosApi";
import { type Service } from "@/types";

export const addService = async (values: Service): Promise<Service> => {
  try {
    const response = await apiClient.post("/services", values);
    return response.data as Service;
  } catch (error: unknown) {
    console.error("Error making new services: ", error);
    throw new Error("Failed to make new services");
  }
};

export const getServices = async (): Promise<Service[]> => {
  try {
    const response = await apiClient.get("/services");
    return response.data as Service[];
  } catch (error: unknown) {
    console.error("Error fetching services: ", error);
    throw new Error("Failed to fetch services");
  }
}