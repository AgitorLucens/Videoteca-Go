import { doGet, doPost } from "./http.service";

export const createUserService = async (user) => {
    return await doPost(
        user,
        '/users',
    );;
};