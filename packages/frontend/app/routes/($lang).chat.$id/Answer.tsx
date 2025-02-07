import { Box } from "react-bootstrap-icons";
import sanitize from "sanitize-html";

export const Answer = () => {
  const renderMainView = () => {
    return (
      <div className="w-full">
        <h1 className="my-8 text-3xl font-light">
          when I use css property on an element in react from emotion react I
          get a typesecript error
        </h1>
        <div className="flex items-center">
          <Box color="black" size={20} />
          <span className="text-lg ml-3">Answer</span>
        </div>
        <div
          className="mt-2 mb-24 font-light"
          dangerouslySetInnerHTML={{
            __html: sanitize("<span>Something here </span>"),
          }}
        />
      </div>
    );
  };

  return (
    <div className="w-full relative grid grid-cols-[2fr_1fr] gap-x-8">
      {renderMainView()}
    </div>
  );
};
