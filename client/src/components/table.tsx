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
    <main className="h-screen w-full pt-6 pr-3 bg-[#e6e6e6]">
      <div className="flex h-full">
        <div className="w-2/12 p-4 flex flex-col gap-3">
          <ul>
            <li className="text-black font-bold text-lg my-4">Forms</li>
          </ul>
          <ul className="grid gap-2">
            <li className="text-black font-bold">+ Create...</li>
            <li className="truncate px-2 py-1 rounded-md bg-[#a259ff] font-semibold text-white">
              Home
            </li>
            {formlistQuery.data &&
              formlistQuery.data.map((form) => {
                return (
                  <li className="truncate px-2 py-1 rounded-md" key={form.id}>
                    {form.name}
                  </li>
                );
              })}
          </ul>
        </div>
        <div className="bg-white shadow-md rounded-t-md  w-full p-3 h-full">
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
