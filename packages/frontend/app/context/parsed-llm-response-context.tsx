import { createContext, useContext, useEffect, useState } from "react";

import {
  parseLLMResponse,
  type IParseLLMResponse,
} from "~/utils/parse-llm-response";

import type React from "react";
import type { ReactNode } from "react";

const defaultValue: IParseLLMResponse = {
  userName: "Traveler",
  preferredLocation: "Unknown",
  budget: "Unknown",
  travelDates: "Unknown",
};

const ParsedLLMResponseContext = createContext<{
  parsedLLMResponse: IParseLLMResponse;
  parseInput: (value: string) => void;
}>({
  parsedLLMResponse: defaultValue,
  parseInput: (v: string) => v,
});

export const useParsedLLMResponse = () => {
  return useContext(ParsedLLMResponseContext);
};

export const ParsedLLMResponseProvider: React.FC<{
  children: ReactNode;
}> = ({ children }) => {
  const [parsedLLMResponse, setParsedLLMResponse] =
    useState<IParseLLMResponse>(defaultValue);

  useEffect(() => {
    const contextResult = localStorage.getItem("@Parsed_llm_response_context");

    if (contextResult) {
      const result = JSON.parse(contextResult);
      setParsedLLMResponse(result);
    }
  }, []);

  const handleParseInput = (input: string) => {
    const result = parseLLMResponse(input);
    localStorage.setItem(
      "@Parsed_llm_response_context",
      JSON.stringify(result)
    );

    setParsedLLMResponse(result);
  };

  return (
    <ParsedLLMResponseContext.Provider
      value={{ parsedLLMResponse, parseInput: handleParseInput }}
    >
      {children}
    </ParsedLLMResponseContext.Provider>
  );
};
