import { useParams } from "@remix-run/react";

import { AppHeader } from "~/shared-components/app-header";
import { Button } from "~/shared-components/button/button";
import { MaxWidthContainer } from "~/shared-components/max-width-container";
import { constructURL, ROUTE_IDS } from "~/utils/route-util";

import { Footer } from "./footer";

import type { MetaFunction } from "@remix-run/node";

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
      <main className="bg-blue-300 mt-20 lg:h-[600px] xl:h-[700px] dark:bg-gray-950">
        <MaxWidthContainer>
          <div className="w-full relative">
            <hr className="my-9 md:mt-14 lg:mt-28 border-1 border-gray-200 dark:border-white/10" />

            <div className="w-full grid grid-rows-[1fr_auto] lg:grid-rows-1 lg:grid-cols-2 lg:gap-y-8 xl:gap-y-11">
              <section className="flex flex-col items-center lg:items-start">
                <h1 className="font-bold tracking-tighter text-gray-900 leading-10 text-4xl mb-4 lg:text-6xl lg:mb-6 dark:text-white">
                  Wander Smart
                </h1>
                <p className="font-light text-center lg:text-left lg:text-xl dark:text-gray-400">
                  Unlock the world with intelligent travel insights. Your
                  personal AI navigator for seamless adventures.
                </p>
                <Button
                  type="link"
                  to={constructURL({
                    routeId: ROUTE_IDS.searchWelcomePage,
                    params: { lang },
                  })}
                  size="large"
                  className="mt-5"
                  colorTheme="white"
                >
                  Try Now
                </Button>
              </section>
            </div>

            <hr className="mt-9 border-1 border-gray-200 dark:border-white/10" />
          </div>
        </MaxWidthContainer>
      </main>
    );
  };

  return (
    <div className="relative h-screen lg:overflow-hidden">
      <AppHeader />
      {renderBody()}
      <Footer />
    </div>
  );
}
