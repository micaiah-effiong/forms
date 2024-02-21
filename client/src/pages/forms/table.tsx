import { QueryClient, QueryClientProvider, useQuery } from "@tanstack/react-query";

const queryClient = new QueryClient();
export const withQueryClientProvider = (Comp: React.FC) => {
    return () => {
        return (
            <QueryClientProvider client={queryClient}>
                <Comp />
            </QueryClientProvider>
        );
    };
};

export const FormList = withQueryClientProvider(function () {
    const formlistQuery = useQuery({
        queryKey: ["forms"],
        async queryFn() {
            const res = await fetch("http://localhost:8000/forms");
            return await res.json();
        }
    });

    console.log(formlistQuery.data);

    return (
        <main className="h-screen w-full bg-white">
            <div className="flex">
                <div className="w-2/12 p-4">
                    <button>create</button>
                    <ul>
                        {formlistQuery.data &&
                            formlistQuery.data.map((form: unknown) => {
                                return <li key={form?._id}> {form?.name}</li>;
                            })}
                    </ul>
                </div>
                <div className="border-2 border-black w-full p-3">
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
                                <td>me</td>
                            </tr>
                        </tbody>
                    </table>
                </div>
            </div>
        </main>
    );
});
