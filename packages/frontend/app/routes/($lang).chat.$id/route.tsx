import type { MetaFunction } from "@remix-run/node";

import { MaxWidthContainer } from "~/shared-components/max-width-container";
import { Messagebox } from "~/shared-components/message-box/message-box";
import { Answer } from "./Answer";

export const meta: MetaFunction = ({ params }) => {
  return [
    { title: `Waka Travel | Chat ${params.id}` },
    { name: "description", content: "Plan your next trip fast and smart" },
  ];
};

export default function Route() {
  const renderAnswers = () => {
    return (
      <section className="relative w-full">
        <ul className="w-full relative">
          <li className="w-full border-b border-b-gray-200">
            <Answer />
          </li>
        </ul>
      </section>
    );
  };

  const renderMessagebox = () => {
    return (
      <footer className="w-full absolute bottom-5 grid grid-cols-[2fr_1fr] gap-x-8">
        <Messagebox size="small" onSendMessage={() => null} />
      </footer>
    );
  };

  return (
    <MaxWidthContainer className="w-4/5 xl:w-[1036px] h-full">
      <article className="w-full h-full relative">
        {renderAnswers()}
        {renderMessagebox()}
      </article>
    </MaxWidthContainer>
  );
}
