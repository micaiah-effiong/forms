import * as api from "./generated";

export const apiConfig = api;

apiConfig.OpenAPI.BASE = import.meta.env.VITE_API_URL;
