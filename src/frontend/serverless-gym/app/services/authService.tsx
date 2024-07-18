export const isAuthenticated = () => {
  const authToken = localStorage.getItem("token") ?? "";

  if (authToken.length > 0) {
    return true;
  } else {
    return false;
  }
};
