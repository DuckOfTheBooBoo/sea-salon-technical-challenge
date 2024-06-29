import apiClient from "./axiosApi";
import { Reservation, ReservationRequest } from "@/types";

export const addReservation = async (values: ReservationRequest): Promise<Reservation> => {
  try {
    const response = await apiClient.post("/reservations", values);
    return response.data.reservation as Reservation;
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

export const deleteReservation = async (reservation: Reservation): Promise<void> => {
  try {
    await apiClient.delete("/reservations/" + reservation.ID);
  } catch (error: unknown) {
    console.error("Error deleting reservation: ", error);
    throw new Error("Failed to delete reservation");
  }
}