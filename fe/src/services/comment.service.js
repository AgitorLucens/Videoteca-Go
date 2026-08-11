import { doGet, doPost, doDelete } from "./http.service";

export const getComments = async (msId, page = 1, limit = 10) => {
  return await doGet(`/movieseries/${msId}/comments?page=${page}&limit=${limit}`);
};

export const createComment = async (msId, comment) => {
  return await doPost({ comment }, `/movieseries/${msId}/comments`);
};

export const deleteComment = async (commentId) => {
  return await doDelete(`/comments/${commentId}`);
};
