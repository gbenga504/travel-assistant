import { useParams } from "@remix-run/react";
import { Globe } from "react-bootstrap-icons";

import { Button } from "~/shared-components/button/button";
import { MaxWidthContainer } from "~/shared-components/max-width-container";
import { constructURL, ROUTE_IDS } from "~/utils/route-utils";

export const AppHeader = () => {
  const { lang } = useParams();

  return (
    <nav className="border-b border-b-gray-200 fixed left-0 top-0 z-30 w-full">
      <MaxWidthContainer className="lg:min-h-20 p-4 flex justify-between items-center">
        <div />
        <div className="flex">
          <Button
            type="link"
            to={constructURL({
              routeId: ROUTE_IDS.pricingPage,
              params: { lang },
            })}
            variant="text"
            size="medium"
            colorTheme="white"
          >
            Pricing
          </Button>
          <Button
            type="link"
            to={constructURL({ routeId: ROUTE_IDS.homePage, params: { lang } })}
            variant="text"
            size="medium"
            colorTheme="white"
            icon={<Globe />}
          >
            English
          </Button>
        </div>
      </MaxWidthContainer>
    </nav>
  );
};
