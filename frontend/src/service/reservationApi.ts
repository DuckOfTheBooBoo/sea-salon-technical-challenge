import { apiClient } from "./axiosApi";
import { ReservationRequest } from "@/types";

export const addReservation = async (values: ReservationRequest) => {
  try {
    const response = await apiClient.post("/reservation", values);
    return response.data;
  } catch (error: unknown) {
    console.error("Error making new reservation: ", error);
    throw new Error("Failed to make new reservation");
  }
};
