import { MaxWidthContainer } from "~/shared-components/max-width-container";

export const Footer = () => {
  return (
    <footer className="absolute bottom-0 left-0 w-full mb-5">
      <MaxWidthContainer>
        <span className="font-medium text-xs mb-3">
          &copy; WakaTravel {new Date().getFullYear()}
        </span>
      </MaxWidthContainer>
    </footer>
  );
};
