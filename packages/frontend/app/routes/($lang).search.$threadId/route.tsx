import { useLoaderData } from "@remix-run/react";
import classNames from "classnames";
import { useState } from "react";

import { useQueryAgent } from "~/hooks/use-query-agent";
import { MaxWidthContainer } from "~/shared-components/max-width-container";
import { Messagebox } from "~/shared-components/message-box/message-box";
import { constructURL, ROUTE_IDS } from "~/utils/route-util";

import { ThreadEntry } from "./ThreadEntry";

import type { LoaderFunction, MetaFunction } from "@remix-run/node";

export const meta: MetaFunction = ({ params }) => {
  return [
    { title: `Waka Travel | Search ${params.id}` },
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

  // TODO: Load all messages for the search and return as JSON
  return Response.json([]);
};

export default function Route() {
  const [isMessageboxGrowing, setIsMessageboxGrowing] = useState(false);
  const data = useLoaderData<typeof loader>();
  const { threadEntries } = useQueryAgent(data);

  const renderThread = () => {
    return (
      <MaxWidthContainer className="w-full overflow-y-scroll xl:w-[1036px] md:px-8">
        <ul className="w-full relative">
          {threadEntries.map((te, index) => (
            <li
              className="w-full border-b border-gray-200 dark:border-white/10"
              key={index}
            >
              <ThreadEntry {...te} />
            </li>
          ))}
        </ul>
      </MaxWidthContainer>
    );
  };

  const renderMessagebox = () => {
    return (
      <footer className="w-full absolute bottom-4 flex justify-center">
        <MaxWidthContainer className="w-full gap-x-8 z-50 grid grid-cols-1 lg:grid-cols-[2fr_1fr] xl:w-[1036px] md:px-8">
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
              onSendMessage={() => null}
              onGrow={(growing) => setIsMessageboxGrowing(growing)}
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
