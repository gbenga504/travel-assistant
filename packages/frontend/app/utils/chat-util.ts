import { nanoid } from "nanoid";

export const encodeChatIdParam = (message: string): string => {
  const formattedMessage = message
    .toLowerCase()
    .substring(0, 25)
    .replace(/\s/g, "-");

  return `${formattedMessage}-${nanoid()}`;
};
