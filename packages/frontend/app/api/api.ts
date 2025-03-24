import axios from "axios";

import { configuration } from "~/utils/configuration";

import { Thread } from "./thread";

export const createApiClient = () => {
  const httpClient = axios.create({
    baseURL: configuration.API_URL,
    timeout: 8000,
  });

  const thread = new Thread(httpClient);

  return { thread };
};

export type ICreateApiClient = ReturnType<typeof createApiClient>;
