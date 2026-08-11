import { doGet } from "./http.service";

export const searchMovies = async (query) => {
  return await doGet(`/search?q=${encodeURIComponent(query)}`);
};