export const encodeChatIdParam = (message: string): string => {
  return message.toLowerCase().substring(0, 25).replace(/\s/g, "-");
};
