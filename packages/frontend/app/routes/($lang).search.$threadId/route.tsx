import { redirect } from "@remix-run/node";
import { useLoaderData } from "@remix-run/react";
import classNames from "classnames";
import { useState } from "react";

import { createApiClient } from "~/api/api";
import { useQueryAgent } from "~/hooks/use-query-agent";
import { MaxWidthContainer } from "~/shared-components/max-width-container";
import { Messagebox } from "~/shared-components/message-box/message-box";
import { constructURL, ROUTE_IDS } from "~/utils/route-util";
import { transformToThreadEntry } from "~/utils/search-util";

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

  const handleSendQuery = (query: string) => {
    queryAgent(query);
    setMessage("");
  };

  const renderThread = () => {
    return (
      <div className="w-full h-full relative overflow-y-scroll flex justify-center">
        <MaxWidthContainer className="w-full h-full xl:w-[772px] md:px-8">
          <ul className="w-full relative pb-36 pt-6">
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
      <footer className="w-full absolute bottom-4 flex justify-center">
        <MaxWidthContainer className="w-full z-50 grid grid-cols-1 xl:w-[772px] md:px-8">
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
    <article className="w-full h-full relative overflow-hidden">
      {renderThread()}
      {renderMessagebox()}
    </article>
  );
}
