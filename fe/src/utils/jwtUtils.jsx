import jwtDecode from "jwt-decode";

// Save token
export const setToken = (token) => {
  localStorage.setItem("token", token);
};

// Get token
export const getToken = () => {
  return localStorage.getItem("token");
};

// Get user info (decoded)
export const getUser = () => {
  const token = getToken();
  if (!token) return null;
  try {
    return jwtDecode(token); // expects token to contain { role: "admin" }
  } catch (err) {
    console.error("Invalid token", err);
    return null;
  }
};

// Logout
export const logout = () => {
  localStorage.removeItem("token");
};