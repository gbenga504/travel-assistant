import { redirect } from "@remix-run/node";
import { useLoaderData } from "@remix-run/react";
import classNames from "classnames";
import { useEffect, useRef, useState } from "react";
import { ClientOnly } from "remix-utils/client-only";

import { createApiClient } from "~/api/api";
import { MapConfigProvider } from "~/context/map-config-context";
import { useQueryAgent } from "~/hooks/use-query-agent";
import { LoadingSpinner } from "~/shared-components/loading-spinner";
import { MaxWidthContainer } from "~/shared-components/max-width-container";
import { Messagebox } from "~/shared-components/message-box/message-box";
import { constructURL, ROUTE_IDS } from "~/utils/route-util";
import { transformToThreadEntry } from "~/utils/search-util";

import { AppHeader } from "./app-header";
import { LazyMap } from "./map/lazy-map";
import { ThreadEntry } from "./thread-entry";

import type { LoaderFunction, MetaFunction } from "@remix-run/node";

export const meta: MetaFunction = () => {
  return [
    { title: `Waka Travel | Search` },
    { name: "description", content: "Plan your next trip fast and smart" },
  ];
};

export const loader: LoaderFunction = async ({ request, params }) => {
  const refererURL = new URL(
    (request.headers.get("Referer") as string) ?? request.url
  );
  const refererPath = refererURL.pathname;
  const searchWelcomePagePath = constructURL({
    routeId: ROUTE_IDS.searchWelcomePage,
    params: { lang: params.lang },
  });

  // If the request was initiated by the search welcome page, then we want to skip loading
  // all messages for the search
  if (refererPath == searchWelcomePagePath) {
    return Response.json([]);
  }

  const api = createApiClient();
  const result = await api.thread.getThread(params.threadId!);

  if (result.data.length === 0) {
    return redirect(
      constructURL({ routeId: ROUTE_IDS.searchWelcomePage, params })
    );
  }

  return Response.json(result.data);
};

export default function Route() {
  const [isMessageboxGrowing, setIsMessageboxGrowing] = useState(false);
  const data = useLoaderData<typeof loader>();
  const { thread, queryAgent } = useQueryAgent(transformToThreadEntry(data));
  const [message, setMessage] = useState("");
  const chatSectionRef = useRef<HTMLDivElement>(null);

  useEffect(
    function scrollToBottom() {
      if (chatSectionRef.current) {
        chatSectionRef.current.scrollTop = chatSectionRef.current.scrollHeight;
      }
    },
    [thread]
  );

  const handleSendQuery = (query: string) => {
    queryAgent(query);
    setMessage("");
  };

  const renderThread = () => {
    return (
      <div
        className="w-full h-full relative overflow-y-scroll flex flex-col justify-center"
        ref={chatSectionRef}
      >
        <MaxWidthContainer className="w-full h-full lg:w-11/12 2xl:w-8/12 md:px-8">
          <ul className="w-full relative pb-10">
            {thread.map((te, index) => (
              <li
                className="w-full border-b border-gray-200 dark:border-white/10"
                key={index}
              >
                <ThreadEntry {...te} />
              </li>
            ))}
          </ul>
        </MaxWidthContainer>
      </div>
    );
  };

  const renderMessagebox = () => {
    return (
      <footer className="w-full flex justify-center mb-2">
        <MaxWidthContainer className="w-full z-50 grid grid-cols-1 lg:w-11/12 2xl:w-8/12 md:px-8">
          <div
            className={classNames(
              "w-full p-2 rounded-full bg-white dark:bg-gray-900",
              {
                "rounded-md": isMessageboxGrowing,
              }
            )}
          >
            <Messagebox
              size="small"
              onSendMessage={handleSendQuery}
              onGrow={(growing) => setIsMessageboxGrowing(growing)}
              value={message}
              onChange={(m) => setMessage(m)}
            />
          </div>
        </MaxWidthContainer>
      </footer>
    );
  };

  return (
    <MapConfigProvider>
      <article className="w-full h-full flex flex-col">
        <AppHeader />
        <main className="flex-1 overflow-hidden w-full grid grid-cols-2">
          <section className="w-full h-full overflow-hidden flex flex-col">
            {renderThread()}
            {renderMessagebox()}
          </section>
          <section className="w-full flex justify-center items-center">
            <ClientOnly fallback={<LoadingSpinner />}>
              {() => <LazyMap />}
            </ClientOnly>
          </section>
        </main>
      </article>
    </MapConfigProvider>
  );
}
