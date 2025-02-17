import { useParams, useNavigate, useLocation } from "@remix-run/react";
import { Globe } from "react-bootstrap-icons";

import { Button } from "~/shared-components/button/button";
import { Dropdown } from "~/shared-components/dropdown";
import { MaxWidthContainer } from "~/shared-components/max-width-container";
import { getLanguagePath, SUPPORTED_LANGUAGES } from "~/utils/language-util";
import { constructURL, ROUTE_IDS } from "~/utils/route-util";

import type { IOption } from "~/shared-components/dropdown";
import type { ISupportedLanguages } from "~/utils/language-util";

const LANGUAGE_MAP: { [key in ISupportedLanguages]: string } = {
  en: "English",
  de: "German",
};

export const AppHeader = () => {
  const { lang } = useParams();
  const navigate = useNavigate();
  const { pathname } = useLocation();

  const handleSelectLanguage = (option: IOption) => {
    navigate(getLanguagePath(pathname, option.value));
  };

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
                {LANGUAGE_MAP[lang as ISupportedLanguages]}
              </Button>
            }
            options={[
              { label: "English", value: SUPPORTED_LANGUAGES.en },
              { label: "German", value: SUPPORTED_LANGUAGES.de },
            ]}
            onSelect={handleSelectLanguage}
          />
        </div>
      </MaxWidthContainer>
    </nav>
  );
};
