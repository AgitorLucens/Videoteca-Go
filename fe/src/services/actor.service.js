import { doGet, doPost, doPut, doDelete } from "./http.service";

export const getActors = async () => {
  return await doGet("/admin/actors");
};

export const createActor = async (data) => {
  return await doPost(data, "/admin/actors");
};

export const updateActor = async (id, data) => {
  return await doPut(data, `/admin/actors/${id}`);
};

export const deleteActor = async (id) => {
  return await doDelete(`/admin/actors/${id}`);
};
