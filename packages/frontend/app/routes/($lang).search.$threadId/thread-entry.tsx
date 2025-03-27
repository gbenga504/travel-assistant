import classNames from "classnames";
import Markdown from "markdown-to-jsx";
import { Box } from "react-bootstrap-icons";

import { LoadingSection } from "~/shared-components/loading-section";

import { markdownOverrides } from "./markdown-overrides";

import type { IThreadEntry } from "~/utils/search-util";

export const ThreadEntry = ({ question, status, answer }: IThreadEntry) => {
  const renderMainView = () => {
    return (
      <div className="w-full">
        <h1 className="my-4 text-xl font-medium">{question}</h1>
        <div className="flex items-center">
          <Box
            className={classNames("text-gray-900  dark:text-white", {
              "animate-bounce": status != "COMPLETED",
            })}
            size={20}
          />
          <span className="ml-3 font-light">Reply</span>
        </div>
        <div className="mt-2 mb-8 font-extralight">
          {status === "PENDING" && <LoadingSection />}
          {status !== "PENDING" && (
            <Markdown
              options={{
                overrides: markdownOverrides,
              }}
            >
              {answer}
            </Markdown>
          )}
        </div>
      </div>
    );
  };

  return (
    <div className="w-full relative grid grid-cols-1">{renderMainView()}</div>
  );
};
