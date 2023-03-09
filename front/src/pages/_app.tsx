import "~/styles/reset.css";
import "~/styles/globals.css";
import "react-loading-skeleton/dist/skeleton.css";

import type { AppProps } from "next/app";

if (process.env.NODE_ENV === "development") {
  // dynamic import でファイルを読み込んで MSW を有効にする
  const MockServer = () =>
    import("~/msw/worker").then((mo) => {
      mo.worker.start();
    });
  MockServer();
}

export default function App({ Component, pageProps }: AppProps) {
  return <Component {...pageProps} />;
}
