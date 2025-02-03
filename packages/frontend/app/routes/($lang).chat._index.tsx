import type { MetaFunction } from "@remix-run/node";
import { useNavigate, useParams } from "@remix-run/react";

import { Messagebox } from "~/shared-components/message-box/message-box";
import { encodeChatIdParam } from "~/utils/chat-util";
import { constructURL, ROUTE_IDS } from "~/utils/route-util";
import { MaxWidthContainer } from "~/shared-components/max-width-container";

export const meta: MetaFunction = () => {
  return [
    { title: "Waka Travel | Chat" },
    { name: "description", content: "Plan your next trip fast and smart" },
  ];
};

export default function Route() {
  const navigate = useNavigate();
  const { lang } = useParams();

  const handleSendMessage = (message: string) => {
    navigate(
      constructURL({
        routeId: ROUTE_IDS.chatPage,
        params: { lang: lang!, id: encodeChatIdParam(message) },
      })
    );
  };

  return (
    <MaxWidthContainer className="w-4/5 lg:w-[640px]">
      <article className="w-full h-screen flex items-center justify-center">
        <div className="w-full">
          <h2 className="text-center text-4xl font-light">
            How can I be of help ?
          </h2>
          <section className="mt-8">
            <Messagebox onSendMessage={handleSendMessage} />
          </section>
        </div>
      </article>
    </MaxWidthContainer>
  );
}
