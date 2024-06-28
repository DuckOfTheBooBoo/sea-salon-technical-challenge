import apiClient from "./axiosApi";
import { Reservation, ReservationRequest } from "@/types";

export const addReservation = async (values: ReservationRequest) => {
  try {
    const response = await apiClient.post("/reservations", values);
    return response.data;
  } catch (error: unknown) {
    console.error("Error making new reservation: ", error);
    throw new Error("Failed to make new reservation");
  }
};

export const getReservationAll = async (): Promise<Reservation[]> => {
  try {
    const response = await apiClient.get("/reservations");
    return response.data;
  } catch (error: unknown) {
    console.error("Error fetching reservations: ", error);
    throw new Error("Failed to fetch reservations");
  }
}