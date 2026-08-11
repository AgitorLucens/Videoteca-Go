import { doGet, doPost, doPut, doDelete } from "./http.service";

export const getGenres = async () => {
  return await doGet("/admin/genres");
};

export const createGenre = async (data) => {
  return await doPost(data, "/admin/genres");
};

export const updateGenre = async (id, data) => {
  return await doPut(data, `/admin/genres/${id}`);
};

export const deleteGenre = async (id) => {
  return await doDelete(`/admin/genres/${id}`);
};
