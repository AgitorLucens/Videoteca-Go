import { doPost } from "./http.service";

export const forgotPassword = async (email) => {
  return await doPost({ email }, "/forgot-password");
};

export const resetPassword = async (email, newPassword) => {
  return await doPost({ email, newPassword }, "/reset-password");
};
