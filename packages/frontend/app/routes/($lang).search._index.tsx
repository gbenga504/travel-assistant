import { useNavigate, useParams } from "@remix-run/react";

import { AppHeader } from "~/shared-components/app-header";
import { MaxWidthContainer } from "~/shared-components/max-width-container";
import { Messagebox } from "~/shared-components/message-box/message-box";
import { constructURL, ROUTE_IDS } from "~/utils/route-util";
import { encodeThreadIdParam } from "~/utils/search-util";

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

  const handleSendMessage = (message: string) => {
    navigate(
      constructURL({
        routeId: ROUTE_IDS.searchPage,
        params: { lang: lang!, id: encodeThreadIdParam(message) },
      }),
      { state: { query: message } }
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
            <Messagebox size="large" onSendMessage={handleSendMessage} />
          </section>
        </div>
      </article>
    </MaxWidthContainer>
  );
}
