import {
  Links,
  Meta,
  Outlet,
  Scripts,
  ScrollRestoration,
} from "@remix-run/react";
import { redirect } from "@remix-run/node";
import type { LinksFunction, LoaderFunction } from "@remix-run/node";

import "./root.css";
import {
  getUserLanguage,
  isSupportedLanguageInPath,
} from "./utils/language-util";

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

    // We get the language which is usually the first pathname. Then we use a heuristic
    // to determine if we should replace this or we just prepend the userLanguage with the fullpath
    // Basically, if the pathname.length == 2, then we replace else we prepend the userLanguage with the fullPath
    const pathname = url.pathname;

    if (pathname === "/") {
      return redirect(`/${userLanguage}`);
    }

    const suppliedLanguage = pathname.split("/")[1];

    // Then the user entered a language but we don't reconginize it
    // we redirect the user to the same page but replace the supplied language with out preferred language
    if (suppliedLanguage.length === 2) {
      return redirect(
        `/${userLanguage}/${pathname.split("/").splice(2).join("/")}`
      );
    }

    // The user did not enter a language.
    // We redirect the user to the same page but add out preferred language to the path
    return redirect(`/${userLanguage}${pathname}`);
  }

  return null;
};

export function Layout({ children }: { children: React.ReactNode }) {
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
        <Scripts />
      </body>
    </html>
  );
}

export default function App() {
  return <Outlet />;
}
