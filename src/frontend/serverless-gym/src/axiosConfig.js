// src/axiosConfig.js
import axios from "axios";

export const api = axios.create({
  baseURL: "https://gcloud-go-7tq7m2dbcq-nw.a.run.app",
  //baseURL: 'http://localhost:5004/order',
});

api.interceptors.request.use((config) => {
  var token = localStorage.getItem("token");

  if (token) {
    config.headers.Authorization = `Bearer ${token}`;
  }
  return config;
});
