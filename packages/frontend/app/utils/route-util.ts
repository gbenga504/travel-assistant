import { generatePath } from "@remix-run/react";

export const ROUTE_IDS = {
  homePage: "homePage",
  searchWelcomePage: "searchWelcomePage",
  searchPage: "searchPage",
  pricingPage: "pricingPage",
} as const;

interface IRoute {
  id: keyof typeof ROUTE_IDS;
  path: string;
}

const routes: IRoute[] = [
  { id: ROUTE_IDS.homePage, path: "/:lang" },
  { id: ROUTE_IDS.searchWelcomePage, path: "/:lang/search" },
  { id: ROUTE_IDS.searchPage, path: "/:lang/search/:id" },
  { id: ROUTE_IDS.pricingPage, path: "/:lang/pricing" },
];

const getPath = ({
  routes,
  routeId,
}: {
  routes: IRoute[];
  routeId: string;
}): string | undefined => {
  if (routes.length === 0) return undefined;

  const [firstRoute, ...restRoutes] = routes;

  if (firstRoute.id === routeId) {
    return firstRoute.path;
  }

  return getPath({ routes: restRoutes, routeId });
};

export const constructURL = ({
  routeId,
  query,
  params,
}: {
  routeId: keyof typeof ROUTE_IDS;
  query?: { [key: string]: string | undefined | null };
  params?: { [key: string]: string | undefined };
}): string => {
  let path = getPath({ routes, routeId });

  if (!path) {
    throw new Error(`Cannot find path with routeId ${routeId}`);
  }

  path = generatePath(path, params);

  if (query) {
    const searchParams = new URLSearchParams();

    Object.keys(query).forEach((key) => {
      if (query[key]) {
        searchParams.append(key, query[key]!);
      }
    });

    path += `?${searchParams.toString()}`;
  }

  return path;
};
