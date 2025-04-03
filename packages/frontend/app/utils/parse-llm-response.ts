export interface IParseLLMResponse {
  userName?: string;
  preferredLocation?: string;
  budget?: string;
  travelDates?: string;
}

export const parseLLMResponse = (response: string): IParseLLMResponse => {
  const html = document.createElement("div");
  html.innerHTML = response;

  const result: IParseLLMResponse = {};

  html.querySelectorAll("span").forEach((span) => {
    const dataType = span.getAttribute("dataType");
    const dataValue = span.getAttribute("dataValue");

    switch (dataType) {
      case "userName":
        result.userName = dataValue ?? undefined;
        break;

      case "location":
        if (span.getAttribute("dataPreference") === "preferred") {
          result.preferredLocation = dataValue ?? undefined;
        }
        break;

      case "budget":
        result.budget = dataValue ?? undefined;
        break;

      case "travelDates":
        result.travelDates = dataValue ?? undefined;
        break;
    }
  });

  return result;
};
