import axios, { type AxiosInstance } from "axios";

let instance: AxiosInstance | undefined;

export function getAxios(): AxiosInstance {
  if (instance === undefined) {
    instance = axios.create({
      baseURL: "http://localhost:8080",
    })
  }
  return instance;
}
