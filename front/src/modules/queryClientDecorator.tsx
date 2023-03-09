import type { ReactRenderer } from "@storybook/react";
import { QueryClient, QueryClientProvider } from "react-query";
import { PartialStoryFn } from "@storybook/csf";

const queryClient = new QueryClient();

export const queryClientDecorator = <T,>(Story: PartialStoryFn<ReactRenderer, T>) => {
  return <QueryClientProvider client={queryClient}>{Story()}</QueryClientProvider>;
};
