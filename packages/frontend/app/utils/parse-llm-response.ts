export interface IExtractUserSettings {
  userName?: string;
}

export const USERNAME_REGEX = /\{\{([^}]+)\}\}/g;

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
      .replace(/\{\{([^}]+)\}\}/g, (_, username) => username)
      .replace(/\*\*(\[\[([^\]]+)\]\])\*\*/g, "$1")
      .replace(/\[\[([^\]]+)\]\]/g, (_, location) => {
        const { name, coordinates } = parseCoordinates(location);

        return `<Location name="${name}" latitude="${coordinates[0]}" longitude="${coordinates[1]}" />`;
      })
      .replace(/\*\*(\[\[([^\]]+)\]\])\*\*/g, "$1")
      .replace(/\*\*([^*]+)\*\*/g, (_, attraction) => {
        const { name, coordinates } = parseCoordinates(attraction);

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
  coordinates: [number, number];
} {
  const [name, lon, lat] = locationString.split(";");

  return {
    name: name.trim(),
    coordinates: [parseFloat(lon), parseFloat(lat)],
  };
}
