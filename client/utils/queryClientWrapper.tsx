import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import React from "react";

export const withQueryClientProvider = (Comp: React.FC, qc: QueryClient) => {
  return () => {
    return (
      <QueryClientProvider client={qc}>
        <Comp />
      </QueryClientProvider>
    );
  };
};
