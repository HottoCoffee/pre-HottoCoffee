import { Button } from "./index";

import type { Meta, StoryObj } from "@storybook/react";
import { queryClientDecorator } from "~/modules/queryClientDecorator";

const meta: Meta<typeof Button> = {
  title: "shared/Button",
  component: Button,
  decorators: [queryClientDecorator],
};

export default meta;
type Story = StoryObj<typeof Button>;

export const Default: Story = {
  args: {
    placeholder: "Placeholder",
    children: "label",
    appearance: "default",
  },
};

export const Success: Story = {
  args: {
    placeholder: "Placeholder",
    children: "label",
    appearance: "success",
  },
};

export const Danger: Story = {
  args: {
    placeholder: "Placeholder",
    children: "label",
    appearance: "danger",
  },
};

export const Labeled: Story = {
  args: {
    placeholder: "Placeholder",
    children: "label",
    appearance: "labeled",
  },
};

export const Disabled: Story = {
  args: {
    placeholder: "Placeholder",
    children: "label",
    appearance: "disabled",
    disabled: true,
  },
};
