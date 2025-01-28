import type { MetaFunction } from "@remix-run/node";

import { Sidebar } from "./sidebar";
import { MaxWidthContainer } from "~/shared-components/max-width-container";
import { Messagebox } from "~/shared-components/message-box/message-box";

export const meta: MetaFunction = () => {
  return [
    { title: "Waka Travel | Chat" },
    { name: "description", content: "Plan your next trip fast and smart" },
  ];
};

export default function Route() {
  const renderEmptyView = () => {
    return (
      <article className="w-full h-screen flex items-center justify-center">
        <div className="w-full">
          <h2 className="text-center text-4xl font-light">
            How can I be of help ?
          </h2>
          <section className="mt-8">
            <Messagebox />
          </section>
        </div>
      </article>
    );
  };

  return (
    <>
      <Sidebar />
      <MaxWidthContainer className="w-4/5 lg:w-[640px]">
        {renderEmptyView()}
      </MaxWidthContainer>
    </>
  );
}
