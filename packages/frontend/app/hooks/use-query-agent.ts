import { useParams } from "@remix-run/react";
import { useEffect, useState } from "react";

import { useApi } from "~/context/api-context";
import { INITIAL_SEARCH_QUERY_KEY } from "~/utils/search-util";

import type { IThreadEntry } from "~/utils/search-util";

export const useQueryAgent = (te: IThreadEntry[]) => {
  const params = useParams<{ threadId: string; lang: string }>();
  const api = useApi();
  const [thread, setThread] = useState<IThreadEntry[]>(te);

  useEffect(() => {
    const query = sessionStorage.getItem(INITIAL_SEARCH_QUERY_KEY);

    // If the request was not triggered from the search index page, then we don't need to
    // trigger an initial search here.
    if (!query) {
      return;
    }

    queryAgent(query);
    sessionStorage.removeItem(INITIAL_SEARCH_QUERY_KEY);
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, []);

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
