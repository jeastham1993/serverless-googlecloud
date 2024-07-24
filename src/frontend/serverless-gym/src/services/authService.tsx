import { NavigateFunction } from "react-router-dom";

export const isAuthenticated = () => {
  return true;
  const authToken = localStorage.getItem("token") ?? "";

  if (authToken.length > 0) {
    return true;
  } else {
    return false;
  }
};

export const logout = (navigate: NavigateFunction) => {
  localStorage.removeItem("token");
  localStorage.removeItem("refreshToken");
  localStorage.removeItem("emailAddress");

  navigate("/login");
};
