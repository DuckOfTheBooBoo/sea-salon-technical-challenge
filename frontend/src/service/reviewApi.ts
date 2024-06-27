import apiClient from "./axiosApi";
import { Review } from "@/types";

export const addReview = async (values: Review): Promise<Review> => {
  try {
    const response = await apiClient.post("/reviews", values);
    return response.data.review as Review;
  } catch (error: unknown) {
    console.error("Error making new review: ", error);
    throw new Error("Failed to make new review");
  }
};

export const getReviews = async (): Promise<Review[]> => {
  try {
    const response = await apiClient.get("/reviews");
    return response.data as Review[];
  } catch (error: unknown) {
    console.error("Error fetching reviews: ", error);
    throw new Error("Failed to fetch reviews");
  }
}