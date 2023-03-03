import { Header } from "./index";

import type { Meta, StoryObj } from "@storybook/react";

const meta: Meta<typeof Header> = {
  title: "shared/Header",
  component: Header,
};

export default meta;
type Story = StoryObj<typeof Header>;

export const HasNoChild: Story = {
  args: {},
};

export const WithChild: Story = {
  args: {
    children: <div style={{ color: "black" }}>Children</div>,
  },
};
