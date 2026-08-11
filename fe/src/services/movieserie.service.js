import { doGet, doPost, doPut, doDelete } from "./http.service";

export const getMovieSeries = async () => {
  return await doGet("/admin/movieseries");
};

export const getMovieSerieById = async (id) => {
  return await doGet(`/movieseries/${id}`);
};

export const createMovieSerie = async (data) => {
  return await doPost(data, "/admin/movieseries");
};

export const updateMovieSerie = async (id, data) => {
  return await doPut(data, `/admin/movieseries/${id}`);
};

export const deleteMovieSerie = async (id) => {
  return await doDelete(`/admin/movieseries/${id}`);
};
