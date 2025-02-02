import type { MetaFunction } from "@remix-run/node";

import { MiniMessagebox } from "~/shared-components/message-box/mini-message-box";

export const meta: MetaFunction = ({ params }) => {
  return [
    { title: `Waka Travel | Chat ${params.id}` },
    { name: "description", content: "Plan your next trip fast and smart" },
  ];
};

export default function Route() {
  return (
    <article className="w-full h-screen relative flex items-center justify-center">
      <section className="absolute bottom-4 w-full">
        <MiniMessagebox />
      </section>
    </article>
  );
}
