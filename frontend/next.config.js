// eslint-disable-next-line @typescript-eslint/no-var-requires
const path = require("path");

/** @type {import('next').NextConfig} */
const nextConfig = {
  reactStrictMode: true,
  assetPrefix: process.env.NEXT_PUBLIC_ENABLE_GH_PAGES ? "/HottoCoffee" : "",
  basePath: process.env.NEXT_PUBLIC_ENABLE_GH_PAGES ? "/HottoCoffee" : "",
  trailingSlash: true,
  images: {
    unoptimized: true,
  },
  experimental: {
    optimizeFonts: true,
  },
  sassOptions: {
    includePaths: [path.join(__dirname, "~/")],
  },
  output: "standalone",
  redirects: async () => {
    return [
      {
        source: "/",
        destination: "/home",
        permanent: true,
      },
    ];
  },
};

module.exports = nextConfig;
