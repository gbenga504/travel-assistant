import { redirect } from "@remix-run/node";
import {
  Links,
  Meta,
  Outlet,
  Scripts,
  ScrollRestoration,
  useLoaderData,
} from "@remix-run/react";

import { createApiClient } from "./api/api";
import { ApiProvider } from "./context/api-context";
import { configuration } from "./utils/configuration";
import {
  getLanguagePath,
  getUserLanguage,
  isSupportedLanguageInPath,
} from "./utils/language-util";

import type { LinksFunction, LoaderFunction } from "@remix-run/node";

import "./root.css";

export const links: LinksFunction = () => [
  { rel: "preconnect", href: "https://fonts.googleapis.com" },
  {
    rel: "preconnect",
    href: "https://fonts.gstatic.com",
    crossOrigin: "anonymous",
  },
  {
    rel: "stylesheet",
    href: "https://fonts.googleapis.com/css2?family=Inter:ital,opsz,wght@0,14..32,100..900;1,14..32,100..900&display=swap",
  },
];

export const loader: LoaderFunction = async ({ request }) => {
  // We want to redirect the user to the right locale website
  const url = new URL(request.url);

  if (!isSupportedLanguageInPath(url.pathname)) {
    const userLanguage = getUserLanguage(request);

    return redirect(getLanguagePath(url.pathname, userLanguage));
  }

  // We need to return environment variables
  return Response.json({
    ENV: {
      API_URL: configuration.API_URL,
    },
  });
};

export function Layout({ children }: { children: React.ReactNode }) {
  const data = useLoaderData<typeof loader>();

  return (
    <html lang="en">
      <head>
        <meta charSet="utf-8" />
        <meta name="viewport" content="width=device-width, initial-scale=1" />
        <Meta />
        <Links />
      </head>
      <body>
        {children}
        <ScrollRestoration />
        <script
          dangerouslySetInnerHTML={{
            __html: `window.ENV = ${JSON.stringify(data.ENV)}`,
          }}
        />
        <Scripts />
      </body>
    </html>
  );
}

export default function App() {
  return (
    <ApiProvider api={createApiClient()}>
      <Outlet />
    </ApiProvider>
  );
}
