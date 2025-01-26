import { Globe } from "react-bootstrap-icons";

import { Button } from "~/shared-components/button/button";
import { MaxWidthContainer } from "~/shared-components/max-width-container";

export const AppHeader = () => {
  return (
    <nav className="border-b border-b-gray-200 fixed left-0 top-0 z-30 w-full">
      <MaxWidthContainer className="lg:min-h-20 p-4 flex justify-between items-center">
        <div />
        <div className="flex">
          <Button type="link" to="/pricing" variant="text" size="large">
            Pricing
          </Button>
          <Button
            type="link"
            to="/lang"
            variant="text"
            size="large"
            icon={<Globe className="mr-1" />}
          >
            English
          </Button>
        </div>
      </MaxWidthContainer>
    </nav>
  );
};
