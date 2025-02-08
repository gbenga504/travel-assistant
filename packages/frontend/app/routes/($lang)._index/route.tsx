import type { MetaFunction } from "@remix-run/node";

import { AppHeader } from "./app-header";
import { MaxWidthContainer } from "~/shared-components/max-width-container";
import { Button } from "~/shared-components/button/button";
import { constructURL, ROUTE_IDS } from "~/utils/route-util";
import { useParams } from "@remix-run/react";
import { Footer } from "./footer";

export const meta: MetaFunction = () => {
  return [
    { title: "Waka Travel | Home" },
    { name: "description", content: "Welcome to waka travel" },
  ];
};

export default function Route() {
  const { lang } = useParams();

  const renderBody = () => {
    return (
      <main className="bg-blue-300 mt-20 lg:h-[600px] xl:h-[700px]">
        <MaxWidthContainer className="grid grid-rows-[1fr_auto] py-9 lg:grid-rows-1 lg:grid-cols-2 lg:gap-y-8 lg:pt-12 xl:gap-y-11">
          <section className="flex flex-col items-center lg:items-start">
            <h1 className="font-bold text-gray-800 leading-10 text-4xl mt-9 mb-4 md:mt-14 lg:text-6xl lg:mt-28 lg:mb-6">
              Wander Smart
            </h1>
            <p className="font-light text-center lg:text-left lg:text-xl">
              Unlock the world with intelligent travel insights. Your personal
              AI navigator for seamless adventures.
            </p>
            <Button
              type="link"
              to={constructURL({
                routeId: ROUTE_IDS.chatWelcomePage,
                params: { lang },
              })}
              size="large"
              className="mt-5"
              colorTheme="white"
            >
              Try Now
            </Button>
          </section>
        </MaxWidthContainer>
      </main>
    );
  };

  return (
    <div className="relative h-full">
      <AppHeader />
      {renderBody()}
      <Footer />
    </div>
  );
}
