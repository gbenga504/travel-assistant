export const SUPPORTED_LANGUAGES = {
  en: "en",
  de: "de",
} as const;

export const SUPPORTED_LANGUAGES_VALUES = Object.values(SUPPORTED_LANGUAGES);
export type ISupportedLanguages = keyof typeof SUPPORTED_LANGUAGES;

export const getUserLanguage = (request: Request): ISupportedLanguages => {
  // Get the accept language from the request
  const acceptsLanguage = request.headers.get("Accept-Language");
  const acceptLanguage = acceptsLanguage?.split(",")[0].split("-")[0];

  // We need to get the preferred langauge based on the languages we support
  const preferredLanguage: ISupportedLanguages =
    SUPPORTED_LANGUAGES[acceptLanguage as ISupportedLanguages] ?? "en";

  return preferredLanguage;
};

export const isSupportedLanguageInPath = (pathname: string) => {
  if (pathname === "/") {
    return false;
  }

  const language = pathname.split("/")[1] as ISupportedLanguages;

  return language in SUPPORTED_LANGUAGES;
};
