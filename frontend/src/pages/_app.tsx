import "~/styles/reset.css";
import "~/styles/globals.css";
import "react-loading-skeleton/dist/skeleton.css";

import type { AppProps } from "next/app";
import { QueryClient, QueryClientProvider } from "react-query";
import { Provider as ToastProvider } from "@radix-ui/react-toast";

const isBrowser = typeof window !== "undefined";

if (isBrowser) {
  if (process.env.NODE_ENV === "development" || Boolean(process.env.ENABLE_GH_PAGES)) {
    const MockServer = () =>
      import("~/msw/worker").then((mo) => {
        mo.worker.start();
      });
    MockServer();
  }
}

const queryClient = new QueryClient();

export default function App({ Component, pageProps }: AppProps) {
  return (
    <QueryClientProvider client={queryClient}>
      <ToastProvider>
        <Component {...pageProps} />
      </ToastProvider>
    </QueryClientProvider>
  );
}
