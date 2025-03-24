import { useLocation, useParams } from "@remix-run/react";
import { useEffect, useRef, useState } from "react";

import { useApi } from "~/context/api-context";

import type { IThreadEntry } from "~/routes/($lang).search.$threadId/ThreadEntry";

export const useQueryAgent = (te: IThreadEntry[]) => {
  const { state } = useLocation();
  const sentInitialRequest = useRef(false);
  const params = useParams<{ threadId: string; lang: string }>();
  const api = useApi();
  const [thread, setThread] = useState<IThreadEntry[]>(te);

  useEffect(() => {
    // If the request was not triggered from the search index page, then we don't need to
    // trigger an initial search here.
    if (!state || !state.query) {
      return;
    }

    if (!sentInitialRequest.current) {
      queryAgent(state.query);
      sentInitialRequest.current = true;
    }
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [state]);

  function queryAgent(query: string) {
    // We want to have some sort of optimistic update so we can give
    // the user some early feedback
    setThread((prev) => [
      ...prev,
      {
        question: query,
        status: "PENDING",
        answer: "",
      },
    ]);

    api.thread.send(
      query,
      params.threadId!,
      function (_err, { done, message }) {
        const status: IThreadEntry["status"] = done
          ? "COMPLETED"
          : "IN_PROGRESS";

        // We need to only reset the last entry
        setThread((prev) => {
          const otherEntries = prev.slice(0, prev.length - 1);
          const lastEntry = prev[prev.length - 1];

          return [
            ...otherEntries,
            {
              question: query,
              status,
              answer: `${lastEntry.answer}${message}`,
            },
          ];
        });
      }
    );
  }

  return { thread, queryAgent };
};
