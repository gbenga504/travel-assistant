import { useNavigate, useParams } from "@remix-run/react";
import { useState } from "react";

import { useParsedLLMResponse } from "~/context/parsed-llm-response-context";
import { AppHeader } from "~/shared-components/app-header";
import { MaxWidthContainer } from "~/shared-components/max-width-container";
import { Messagebox } from "~/shared-components/message-box/message-box";
import { constructURL, ROUTE_IDS } from "~/utils/route-util";
import {
  encodeThreadIdParam,
  INITIAL_SEARCH_QUERY_KEY,
} from "~/utils/search-util";

import type { MetaFunction } from "@remix-run/node";

export const meta: MetaFunction = () => {
  return [
    { title: "Waka Travel | Search" },
    { name: "description", content: "Plan your next trip fast and smart" },
  ];
};

export default function Route() {
  const navigate = useNavigate();
  const { lang } = useParams();
  const [message, setMessage] = useState("");
  const { clear: clearParsedLLMResponse } = useParsedLLMResponse();

  const handleSendQuery = (query: string) => {
    sessionStorage.setItem(INITIAL_SEARCH_QUERY_KEY, query);
    clearParsedLLMResponse();

    navigate(
      constructURL({
        routeId: ROUTE_IDS.searchPage,
        params: { lang: lang!, id: encodeThreadIdParam(query) },
      })
    );
  };

  return (
    <MaxWidthContainer className="w-full xl:w-[640px]">
      <div className="w-full block lg:hidden">
        <AppHeader />
      </div>

      <article className="w-full h-screen flex items-center justify-center">
        <div className="w-full">
          <h2 className="text-center text-3xl font-light lg:text-4xl">
            How can I be of help ?
          </h2>
          <section className="mt-4 lg:mt-8">
            <Messagebox
              size="large"
              onSendMessage={handleSendQuery}
              value={message}
              onChange={(m) => setMessage(m)}
            />
          </section>
        </div>
      </article>
    </MaxWidthContainer>
  );
}
