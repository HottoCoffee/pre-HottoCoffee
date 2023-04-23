import "~/styles/reset.css";
import "~/styles/globals.css";
import "react-loading-skeleton/dist/skeleton.css";

import type { AppProps } from "next/app";
import { QueryClient, QueryClientProvider } from "react-query";
import { Provider as ToastProvider } from "@radix-ui/react-toast";

if (process.env.NODE_ENV === "development" || Boolean(process.env.SHOULD_ENABLE_MOCK)) {
  // dynamic import でファイルを読み込んで MSW を有効にする
  const MockServer = () =>
    import("~/msw/worker").then((mo) => {
      mo.worker.start();
    });
  MockServer();
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
