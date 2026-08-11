import { doGet, doPut } from "./http.service";

export const getProfile = async () => {
  return await doGet("/user/profile");
};

export const updateEmail = async (email) => {
  return await doPut({ email }, "/user/email");
};

export const updatePassword = async (currentPassword, newPassword) => {
  return await doPut({ currentPassword, newPassword }, "/user/password");
};

export const updateUsername = async (username) => {
  return await doPut({ username }, "/user/username");
};

export const updatePhoto = async (photo) => {
  return await doPut({ photo }, "/user/picture");
};
