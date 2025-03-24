import { nanoid } from "nanoid";

export const INITIAL_SEARCH_QUERY_KEY = "@INITIAL_SEARCH_QUERY_KEY";

export const encodeThreadIdParam = (message: string): string => {
  const formattedMessage = message
    .toLowerCase()
    .substring(0, 25)
    .replace(/\W/gi, "-");

  return `${formattedMessage}-${nanoid()}`;
};
