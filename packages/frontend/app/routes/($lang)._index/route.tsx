import type { MetaFunction } from "@remix-run/node";

import { AppHeader } from "./app-header";
import { MaxWidthContainer } from "../../shared-components/max-width-container";

export const meta: MetaFunction = () => {
  return [
    { title: "Waka Travel | Home" },
    { name: "description", content: "Welcome to waka travel" },
  ];
};

export default function Index() {
  const renderBody = () => {
    return (
      <main className="bg-blue-300 lg:mt-20 lg:max-h-[600px]">
        <MaxWidthContainer className="grid grid-cols-2 gap-y-11 pt-12">
          <section>
            <h1 className="font-bold text-6xl text-gray-800 mt-28 mb-6 leading-10">
              Wander Smart
            </h1>
            <p className="text-xl font-light">
              Unlock the world with intelligent travel insights. Your personal
              AI navigator for seamless adventures.
            </p>
            <div></div>
          </section>
        </MaxWidthContainer>
      </main>
    );
  };

  return (
    <div>
      <AppHeader />
      {renderBody()}
    </div>
  );
}
