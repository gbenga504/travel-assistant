import type { MetaFunction } from "@remix-run/node";

import { MaxWidthContainer } from "~/shared-components/max-width-container";
import { MiniMessagebox } from "~/shared-components/message-box/mini-message-box";

export const meta: MetaFunction = ({ params }) => {
  return [
    { title: `Waka Travel | Chat ${params.id}` },
    { name: "description", content: "Plan your next trip fast and smart" },
  ];
};

export default function Route() {
  return (
    <MaxWidthContainer className="w-4/5 xl:w-[1036px] h-full">
      <article className="w-full h-full relative grid grid-cols-[2fr_1fr] gap-x-8">
        <section className="relative w-full">
          <div className="sticky top-[100vh] pb-5 w-full">
            <MiniMessagebox />
          </div>
        </section>
      </article>
    </MaxWidthContainer>
  );
}
