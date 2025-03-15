import {
  fetchEventSource,
  EventStreamContentType,
} from "@microsoft/fetch-event-source";

import { FatalError } from "~/errors/fatal-error";

import type { AxiosInstance } from "axios";

interface IMessage {
  textCompleted: boolean;
  text: string;
}

export class Ask {
  constructor(private httpClient: AxiosInstance) {}

  async send(
    query: string,
    callback: (err: Error | null, message: IMessage) => void
  ) {
    const ctrl = new AbortController();

    fetchEventSource(`${window.ENV.API_URL}/ask`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({ query }),
      signal: ctrl.signal,
      // Since we don't have last event ID implemented, we want to keep
      // the connection open even when the document is not visible
      openWhenHidden: true,

      async onopen(response) {
        const { ok, headers, status } = response;

        if (ok && headers.get("content-type") === EventStreamContentType) {
          return; // everything's good
        } else if (status >= 400 && status < 500 && status !== 429) {
          // client-side errors are usually non-retriable:
          throw new FatalError();
        } else {
          // We don't have retrieable errors at the moment but probably we want to
          // throw one in future
          return;
        }
      },

      onmessage(msg) {
        // if the server emits an error message, throw an exception
        // so it gets handled by the onerror callback below:
        if (msg.event === "FatalError") {
          throw new FatalError();
        }

        // When we receive this event, the client needs to end the connection
        if (msg.event === "EndStream") {
          ctrl.abort();
        }

        // TODO: Properly type data maybe using zod
        const data = JSON.parse(msg.data);
        callback(null, data);
      },

      onclose() {
        // if the server closes the connection expectedly or unexpectedly
        // we want to retry by throwing a retriable error in future
        ctrl.abort();
      },

      onerror(err) {
        if (err instanceof FatalError) {
          callback(err, { text: "", textCompleted: false });
        } else {
          // do nothing to automatically retry. You can also
          // return a specific retry interval here.
          // We don't want to automatically retry when there is an error so we just abort for now
          ctrl.abort();
        }
      },
    });
  }
}
