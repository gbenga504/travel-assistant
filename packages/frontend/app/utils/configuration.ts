import { isDOM } from "./dom";

export const configuration = {
  API_URL: isDOM ? window.ENV.API_URL : (process.env.API_URL as string),
};
