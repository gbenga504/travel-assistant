import axios from "axios";

import { configuration } from "~/utils/configuration";

import { Ask } from "./ask";

export const createApiClient = () => {
  const httpClient = axios.create({
    baseURL: configuration.API_URL,
    timeout: 8000,
  });

  const ask = new Ask(httpClient);

  return { ask };
};

export type ICreateApiClient = ReturnType<typeof createApiClient>;
