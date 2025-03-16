import { Box } from "react-bootstrap-icons";

import { LoadingSection } from "~/shared-components/loading-section";

export interface IThreadEntry {
  question: string;
  loading: boolean;
}

export const ThreadEntry = ({ question, loading }: IThreadEntry) => {
  const renderMainView = () => {
    return (
      <div className="w-full">
        <h1 className="my-8 text-3xl font-light">{question}</h1>
        <div className="flex items-center">
          <Box
            className="text-gray-900 animate-bounce  dark:text-white"
            size={20}
          />
          <span className="text-lg ml-3">Answer</span>
        </div>
        <div className="mt-2 mb-24 font-light">
          {loading && <LoadingSection />}
        </div>
      </div>
    );
  };

  return (
    <div className="w-full relative gap-x-8 grid grid-cols-1 lg:grid-cols-[2fr_1fr]">
      {renderMainView()}
    </div>
  );
};
