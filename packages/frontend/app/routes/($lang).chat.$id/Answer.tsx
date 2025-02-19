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
          <Box className="text-gray-900 dark:text-white" size={20} />
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
    <div className="w-full relative gap-x-8 grid grid-cols-1 lg:grid-cols-[2fr_1fr]">
      {renderMainView()}
    </div>
  );
};
