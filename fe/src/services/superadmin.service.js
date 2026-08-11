import { doGet, doPost, doPut } from "./http.service";

export const getAdmins = async () => {
  return await doGet("/superadmin/admins");
};

export const createAdmin = async (data) => {
  return await doPost(data, "/superadmin/admins");
};

export const updateAdminPassword = async (userId, newPassword) => {
  return await doPut({ userId, newPassword }, "/superadmin/admins/password");
};
