import React from "react";
import ReactDOM from "react-dom/client";
import "./tw.ts";
import { FormList } from "@/components/table";
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";

ReactDOM.createRoot(document.getElementById("root")!).render(
  <React.StrictMode>
    <QueryClientProvider client={new QueryClient()}>
      <FormList />
    </QueryClientProvider>
  </React.StrictMode>,
);
