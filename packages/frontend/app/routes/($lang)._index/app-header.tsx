import { useParams } from "@remix-run/react";
import { Globe } from "react-bootstrap-icons";

import { Button } from "~/shared-components/button/button";
import { Dropdown } from "~/shared-components/dropdown";
import { MaxWidthContainer } from "~/shared-components/max-width-container";
import { SUPPORTED_LANGUAGES } from "~/utils/language-util";
import { constructURL, ROUTE_IDS } from "~/utils/route-util";

export const AppHeader = () => {
  const { lang } = useParams();

  return (
    <nav className="border-b border-gray-200 bg-white fixed left-0 top-0 z-30 w-full dark:bg-gray-950 dark:border-white/10">
      <MaxWidthContainer className="h-20 p-4 flex justify-between items-center">
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
          <Dropdown
            trigger={
              <Button
                type="button"
                variant="text"
                size="medium"
                colorTheme="white"
                icon={<Globe />}
              >
                English
              </Button>
            }
            options={[
              { label: "English", value: SUPPORTED_LANGUAGES.en },
              { label: "German", value: SUPPORTED_LANGUAGES.de },
            ]}
            onSelect={(opt) => console.log(opt)}
          />
        </div>
      </MaxWidthContainer>
    </nav>
  );
};
