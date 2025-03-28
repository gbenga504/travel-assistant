import { lazy, Suspense } from "react";

import { LoadingSpinner } from "~/shared-components/loading-spinner";

const Map = lazy(() => import("./map"));

export function LazyMap() {
  return (
    <Suspense fallback={<LoadingSpinner />}>
      <Map />
    </Suspense>
  );
}
