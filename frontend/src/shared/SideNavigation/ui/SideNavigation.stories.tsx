import { SideNavigation } from "./index";

import type { Meta, StoryObj } from "@storybook/react";

const meta: Meta<typeof SideNavigation> = {
  title: "shared/SideNavigation",
  component: SideNavigation,
};

export default meta;
type Story = StoryObj<typeof SideNavigation>;

export const HasNoChild: Story = {
  args: {},
};
