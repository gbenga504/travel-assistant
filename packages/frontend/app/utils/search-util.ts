import { nanoid } from "nanoid";

import type { IAPIThread } from "./api.types";

export interface IThreadEntry {
  question: string;
  status: "PENDING" | "IN_PROGRESS" | "COMPLETED";
  answer: string;
}

export const INITIAL_SEARCH_QUERY_KEY = "@INITIAL_SEARCH_QUERY_KEY";

export const encodeThreadIdParam = (message: string): string => {
  const formattedMessage = message
    .toLowerCase()
    .substring(0, 25)
    .replace(/\W/gi, "-");

  return `${formattedMessage}-${nanoid()}`;
};

export const transformToThreadEntry = (
  apiThread: IAPIThread[]
): IThreadEntry[] => {
  return apiThread.map((at) => {
    return {
      question: at.question,
      answer: at.answer,
      status: "COMPLETED",
    };
  });
};
