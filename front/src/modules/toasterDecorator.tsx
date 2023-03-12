import type { ReactRenderer } from "@storybook/react";
import { PartialStoryFn } from "@storybook/csf";
import { Provider } from "@radix-ui/react-toast";

export const toasterDecorator = <T,>(Story: PartialStoryFn<ReactRenderer, T>) => {
  return <Provider>{Story()}</Provider>;
};
