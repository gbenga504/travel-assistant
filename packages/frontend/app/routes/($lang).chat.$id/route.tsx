import classNames from "classnames";
import { useState } from "react";

import { MaxWidthContainer } from "~/shared-components/max-width-container";
import { Messagebox } from "~/shared-components/message-box/message-box";

import { Answer } from "./Answer";

import type { MetaFunction } from "@remix-run/node";

export const meta: MetaFunction = ({ params }) => {
  return [
    { title: `Waka Travel | Chat ${params.id}` },
    { name: "description", content: "Plan your next trip fast and smart" },
  ];
};

export default function Route() {
  const [isMessageboxGrowing, setIsMessageboxGrowing] = useState(false);

  const renderAnswers = () => {
    return (
      <MaxWidthContainer className="w-full overflow-y-scroll xl:w-[1036px] md:px-8">
        <ul className="w-full relative">
          <li className="w-full border-b border-gray-200 dark:border-white/10">
            <Answer />
          </li>
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
      {renderAnswers()}
      {renderMessagebox()}
    </article>
  );
}
