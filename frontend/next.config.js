// eslint-disable-next-line @typescript-eslint/no-var-requires
const path = require("path");

/** @type {import('next').NextConfig} */
const nextConfig = {
  reactStrictMode: true,
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
};

module.exports = nextConfig;
