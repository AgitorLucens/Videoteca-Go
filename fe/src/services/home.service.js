import { doGet } from "./http.service";

export const homeService = async () => {
  return await doGet("/movieseries");
};
