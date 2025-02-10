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

export const isSupportedLanguageInPath = (pathname: string): boolean => {
  if (pathname === "/") {
    return false;
  }

  const language = pathname.split("/")[1] as ISupportedLanguages;

  return language in SUPPORTED_LANGUAGES;
};

export const getLanguagePath = (
  pathname: string,
  userLanguage: string
): string => {
  if (pathname === "/") {
    return `/${userLanguage}`;
  }

  const suppliedLanguage = pathname.split("/")[1];

  // Then the user entered a language but we don't reconginize it
  // we redirect the user to the same page but replace the supplied language with out preferred language
  if (suppliedLanguage.length === 2) {
    const pathWithoutLanguage = pathname.split("/").splice(2).join("/");

    return `/${userLanguage}${
      pathWithoutLanguage.length === 0 ? "" : "/"
    }${pathWithoutLanguage}`;
  }

  // The user did not enter a language.
  // We redirect the user to the same page but add out preferred language to the path
  return `/${userLanguage}${pathname}`;
};
