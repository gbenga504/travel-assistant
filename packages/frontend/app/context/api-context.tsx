import { createContext, useContext } from "react";

import { createApiClient } from "../api/api";

import type { ICreateApiClient } from "../api/api";
import type React from "react";
import type { ReactNode } from "react";

const ApiContext = createContext<ICreateApiClient>(createApiClient());

export const useApi = () => {
  return useContext(ApiContext);
};

export const ApiProvider: React.FC<{
  children: ReactNode;
  api: ICreateApiClient;
}> = ({ children, api }) => {
  return <ApiContext.Provider value={api}>{children}</ApiContext.Provider>;
};
