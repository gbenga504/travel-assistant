import {
  fetchEventSource,
  EventStreamContentType,
} from "@microsoft/fetch-event-source";

import { FatalError } from "~/errors/fatal-error";

import type { AxiosInstance } from "axios";

export class Thread {
  constructor(private httpClient: AxiosInstance) {}

  async send(
    query: string,
    threadId: string,
    callback: (
      err: Error | null,
      { done, message }: { done: boolean; message: string }
    ) => void
  ) {
    const ctrl = new AbortController();

    fetchEventSource(`${window.ENV.API_URL}/threads/ask`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({ query, threadId }),
      signal: ctrl.signal,
      // Since we don't have last event ID implemented, we want to keep
      // the connection open even when the document is not visible
      openWhenHidden: true,

      async onopen(response) {
        const { ok, headers } = response;

        if (ok && headers.get("content-type") === EventStreamContentType) {
          return; // everything's good
        }

        // all other errors are usually non-retriable
        throw new FatalError();
      },

      onmessage(msg) {
        // if the server emits an error message, throw an exception
        // so it gets handled by the onerror callback below:
        if (msg.event === "fatal_error") {
          throw new FatalError();
        }

        // When we receive this event, the client needs to end the connection
        if (msg.event === "end_stream") {
          const data = JSON.parse(msg.data);
          callback(null, { message: data.message, done: true });

          return ctrl.abort();
        }

        // TODO: Properly type data maybe using zod
        const data = JSON.parse(msg.data);

        callback(null, { message: data.message, done: false });
      },

      onclose() {
        // if the server closes the connection expectedly or unexpectedly
        // we want to retry by throwing a retriable error in future
        // For now, we throw a FatalError so we skip retrying
        throw new FatalError();
      },

      onerror(err) {
        // For now, we don't want to retry any errors so we throw to stop the operation
        // TODO: Handle error cases properly
        throw err;
      },
    });
  }

  async getThread(id: string) {
    const { data } = await this.httpClient.get(`/threads/${id}`);

    return data;
  }
}
