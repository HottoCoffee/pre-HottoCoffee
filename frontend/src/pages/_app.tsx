import "~/styles/reset.css";
import "~/styles/globals.css";
import "react-loading-skeleton/dist/skeleton.css";

import type { AppProps } from "next/app";
import { QueryClient, QueryClientProvider } from "react-query";
import { Provider as ToastProvider } from "@radix-ui/react-toast";

const isBrowser = typeof window !== "undefined";
const isGHPages = Boolean(process.env.NEXT_PUBLIC_ENABLE_GH_PAGES);

if (isBrowser) {
  if (process.env.NODE_ENV === "development" || isGHPages) {
    const MockServer = () =>
      import("~/msw/worker").then((mo) => {
        mo.worker.start({
          serviceWorker: {
            url: `${isGHPages ? "/HottoCoffee" : ""}/mockServiceWorker.js`,
          },
        });
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
