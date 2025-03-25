import classNames from "classnames";
import { Box } from "react-bootstrap-icons";

import { LoadingSection } from "~/shared-components/loading-section";

import type { IThreadEntry } from "~/utils/search-util";

export const ThreadEntry = ({ question, status, answer }: IThreadEntry) => {
  const renderMainView = () => {
    return (
      <div className="w-full">
        <h1 className="my-8 text-3xl font-light">{question}</h1>
        <div className="flex items-center">
          <Box
            className={classNames("text-gray-900  dark:text-white", {
              "animate-bounce": status != "COMPLETED",
            })}
            size={20}
          />
          <span className="text-lg ml-3">Answer</span>
        </div>
        <div className="mt-2 mb-24 font-light">
          {status === "PENDING" && <LoadingSection />}
          {status !== "PENDING" && (
            <p dangerouslySetInnerHTML={{ __html: answer }} />
          )}
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
