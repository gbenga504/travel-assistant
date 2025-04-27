export interface IExtractUserSettings {
  userName?: string;
}

export const USERNAME_REGEX = /\{\{([^}]+)\}\}/g;
export const LOCATION_REGEX = /\[\[([^\]]+)\]\]/g;
export const ATTRACTION_REGEX = /\*\*([^*]+)\*\*/g;
export const BOLD_LOCATION_REGEX = /\*\*(\[\[([^\]]+)\]\])\*\*/g;
export const BOLD_ATTRACTION_REGEX = /\*\*(\[\[([^\]]+)\]\])\*\*/g;

export const extractUserSettings = (response: string): IExtractUserSettings => {
  const usernames = response.match(USERNAME_REGEX) || [];
  const userName = usernames.map((match) => match.slice(2, -2))[0];

  const result: IExtractUserSettings = {};

  if (userName) {
    result.userName = userName;
  }

  return result;
};

export const parseLLMResponse = (response: string): string => {
  const formattedResponse = response.split("/n").reduce((acc, line) => {
    const formattedText = line
      .replace(USERNAME_REGEX, (_, username) => username)
      .replace(BOLD_LOCATION_REGEX, "$1")
      .replace(LOCATION_REGEX, (_, location) => {
        const { name, coordinates } = parseCoordinates(location);

        if (!coordinates) {
          return `<Location name="${name}" />`;
        }

        return `<Location name="${name}" latitude="${coordinates[0]}" longitude="${coordinates[1]}" />`;
      })
      .replace(BOLD_ATTRACTION_REGEX, "$1")
      .replace(ATTRACTION_REGEX, (_, attraction) => {
        const { name, coordinates } = parseCoordinates(attraction);

        if (!coordinates) {
          return `<Attraction name="${name}" />`;
        }

        return `<Attraction name="${name}" latitude="${coordinates[0]}" longitude="${coordinates[1]}" />`;
      });

    return `${acc}${formattedText}`;
  }, "");

  return formattedResponse
    .replace(/\n{2,}/g, "\n")
    .replace(/\n/g, "<p style='margin-top: 10px' />");
};

function parseCoordinates(locationString: string): {
  name: string;
  coordinates: [number, number] | null;
} {
  const [name, lon, lat] = locationString.split(";");

  return {
    name: name.trim(),
    coordinates: lon && lat ? [parseFloat(lat), parseFloat(lon)] : null,
  };
}
