import axios, { type AxiosInstance } from "axios";

let instance: AxiosInstance | undefined;

export function getAxios(): AxiosInstance {
  if (instance === undefined) {
    instance = axios.create({
      baseURL: import.meta.env.VITE_API_BASEURL,
    })
  }
  return instance;
}
