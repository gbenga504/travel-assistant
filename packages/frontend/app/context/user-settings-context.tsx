import { createContext, useContext, useEffect, useState } from "react";

import {
  extractUserSettings,
  type IExtractUserSettings,
} from "~/utils/parse-llm-response";

import type React from "react";
import type { ReactNode } from "react";

type IUserSettings = Required<Pick<IExtractUserSettings, "userName">>;

const defaultValue: IUserSettings = {
  userName: "Traveler",
};

const UserSettingsContext = createContext<{
  userSettings: IUserSettings;
  parseInput: (value: string) => void;
  clear: () => void;
}>({
  userSettings: defaultValue,
  parseInput: (v: string) => v,
  clear: () => null,
});

export const useUserSettings = () => {
  return useContext(UserSettingsContext);
};

const STORAGE_KEY = "@USER_SETTINGS";

export const UserSettingsProvider: React.FC<{
  children: ReactNode;
}> = ({ children }) => {
  const [userSettings, setUserSettings] = useState<IUserSettings>(defaultValue);

  useEffect(() => {
    const contextResult = localStorage.getItem(STORAGE_KEY);

    if (contextResult) {
      const result = JSON.parse(contextResult);
      setUserSettings((prev) => ({ ...prev, ...result }));
    }
  }, []);

  const clear = () => {
    localStorage.removeItem(STORAGE_KEY);
    setUserSettings(defaultValue);
  };

  const handleParseInput = (input: string) => {
    const result = { ...userSettings, ...extractUserSettings(input) };
    localStorage.setItem(STORAGE_KEY, JSON.stringify(result));

    setUserSettings((prev) => ({ ...prev, ...result }));
  };

  return (
    <UserSettingsContext.Provider
      value={{ userSettings, parseInput: handleParseInput, clear }}
    >
      {children}
    </UserSettingsContext.Provider>
  );
};
