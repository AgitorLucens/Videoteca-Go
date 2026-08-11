import { doPost } from "./http.service";

export const rateMovie = async (id, rating) => {
  return await doPost({ rating }, `/movieseries/${id}/rate`);
};
