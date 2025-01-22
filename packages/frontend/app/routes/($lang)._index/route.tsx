import type { MetaFunction } from "@remix-run/node";

import { AppHeader } from "./app-header";

export const meta: MetaFunction = () => {
  return [
    { title: "Waka Travel | Home" },
    { name: "description", content: "Welcome to waka travel" },
  ];
};

export default function Index() {
  return (
    <div>
      <AppHeader />
    </div>
  );
}
