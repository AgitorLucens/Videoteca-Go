import { doGet, doPost } from "./http.service";

export const loginService = async (user) => {
	return await doPost(
		user,
		'/login',
	);;
};