import type { Preview } from "@storybook/react";
import { initialize, mswDecorator } from "msw-storybook-addon";
import "react-loading-skeleton/dist/skeleton.css";
import MockDate from "mockdate";

MockDate.set(new Date("2022-12-15T00:00:00"));

initialize();

const preview: Preview = {
  parameters: {
    backgrounds: {
      default: "light",
    },
    actions: { argTypesRegex: "^on[A-Z].*" },
    controls: {
      matchers: {
        color: /(background|color)$/i,
        date: /Date$/,
      },
    },
  },
};

export const decorators = [mswDecorator];

export default preview;
