import axios from 'axios';

const apiInstance = axios.create({
  baseURL: 'http://localhost:8080',
});

apiInstance.interceptors.request.use(
  (config) => {
    const userString = localStorage.getItem('user');
    if (userString) {
      try {
        const user = JSON.parse(userString);
        if (user?.token) {
          config.headers.Authorization = `Bearer ${user.token}`;
        }
      } catch (e) {
        console.error("Error parsing user from localStorage:", e);
      }
    }
    return config;
  },
  (error) => {
    return Promise.reject(error);
  }
);

apiInstance.interceptors.response.use(
  (response) => response,
  (error) => {
    if (error.response?.status === 401) {
      localStorage.removeItem("user");
      localStorage.removeItem("token");
      window.location.href = "/";
    }
    return Promise.reject(error);
  }
);

export const doPost = async (payload, path) => {
  const response = await apiInstance.post(path, payload);
  return response.data;
};

export const doGet = async (path) => {
  const response = await apiInstance.get(path);
  return response.data;
};

export const doPut = async (payload, path) => {
  const response = await apiInstance.put(path, payload);
  return response.data;
};

export const doDelete = async (path) => {
  const response = await apiInstance.delete(path);
  return response.data;
};
