import { createContext, useContext, useEffect, useState } from "react";

import {
  parseLLMResponse,
  type IParseLLMResponse,
} from "~/utils/parse-llm-response";

import type React from "react";
import type { ReactNode } from "react";

type RParsedLLMResponse = Required<
  Pick<
    IParseLLMResponse,
    "budget" | "preferredLocation" | "travelDates" | "userName"
  >
>;

const defaultValue: RParsedLLMResponse = {
  userName: "Traveler",
  preferredLocation: "Unknown",
  budget: "Unknown",
  travelDates: "Unknown",
};

const ParsedLLMResponseContext = createContext<{
  parsedLLMResponse: RParsedLLMResponse;
  parseInput: (value: string) => void;
  clear: () => void;
}>({
  parsedLLMResponse: defaultValue,
  parseInput: (v: string) => v,
  clear: () => null,
});

export const useParsedLLMResponse = () => {
  return useContext(ParsedLLMResponseContext);
};

const STORAGE_KEY = "@PARSED_LLM_RESPONSE_CONTEXT";

export const ParsedLLMResponseProvider: React.FC<{
  children: ReactNode;
}> = ({ children }) => {
  const [parsedLLMResponse, setParsedLLMResponse] =
    useState<RParsedLLMResponse>(defaultValue);

  useEffect(() => {
    const contextResult = localStorage.getItem(STORAGE_KEY);

    if (contextResult) {
      const result = JSON.parse(contextResult);
      setParsedLLMResponse((prev) => ({ ...prev, ...result }));
    }
  }, []);

  const clear = () => {
    localStorage.removeItem(STORAGE_KEY);
    setParsedLLMResponse(defaultValue);
  };

  const handleParseInput = (input: string) => {
    const result = parseLLMResponse(input);
    localStorage.setItem(STORAGE_KEY, JSON.stringify(result));

    setParsedLLMResponse((prev) => ({ ...prev, ...result }));
  };

  return (
    <ParsedLLMResponseContext.Provider
      value={{ parsedLLMResponse, parseInput: handleParseInput, clear }}
    >
      {children}
    </ParsedLLMResponseContext.Provider>
  );
};
