import { DatePickInput } from "./index";

import type { Meta, StoryObj } from "@storybook/react";

const meta: Meta<typeof DatePickInput> = {
  title: "shared/DatePicker/DatePickInput",
  component: DatePickInput,
};

export default meta;
type Story = StoryObj<typeof DatePickInput>;

export const Normal: Story = {
  args: {},
};
