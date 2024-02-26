import { apiConfig } from "@/sdk";
import {
  QueryClient,
  QueryClientProvider,
  useQuery,
} from "@tanstack/react-query";
import React from "react";

const queryClient = new QueryClient();
const withQueryClientProvider = (Comp: React.FC) => {
  return () => {
    return (
      <QueryClientProvider client={queryClient}>
        <Comp />
      </QueryClientProvider>
    );
  };
};

console.log("url", import.meta.env);

export const FormList = withQueryClientProvider(function () {
  const formlistQuery = useQuery({
    queryKey: ["forms"],
    async queryFn() {
      return apiConfig.FormService.getForms();
    },
  });

  console.log(formlistQuery.data);

  return (
    <main className="h-screen w-full bg-white">
      <div className="flex h-full">
        <div className="w-2/12 p-4">
          <button>create</button>
          <ul>
            {formlistQuery.data &&
              formlistQuery.data.map((form) => {
                return <li key={form.id}> {form.name}</li>;
              })}
          </ul>
        </div>
        <div className="border-2 border-black w-full p-3 h-full">
          <table>
            <thead>
              <tr>
                <th>#No</th>
                <th>Name</th>
              </tr>
            </thead>
            <tbody>
              <tr>
                <td>1</td>
                <td>askfns</td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>
    </main>
  );
});
