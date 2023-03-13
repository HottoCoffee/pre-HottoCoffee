import type { StorybookConfig } from "@storybook/nextjs";
import path from "path";
const config: StorybookConfig = {
  stories: ["../src/**/*.mdx", "../src/**/*.stories.@(js|jsx|ts|tsx)"],
  addons: ["@storybook/addon-links", "@storybook/addon-essentials", "@storybook/addon-interactions", "@storybook/addon-a11y", "@storybook/addon-mdx-gfm"],
  framework: {
    name: "@storybook/nextjs",
    options: {}
  },
  webpackFinal: async config => {
    config!.resolve!.alias = {
      ...config!.resolve!.alias,
      "~": path.resolve(__dirname, "../src")
    };
    return config;
  },
  staticDirs: ["../public", "../src"],
  docs: {
    autodocs: "tag"
  }
};
export default config;