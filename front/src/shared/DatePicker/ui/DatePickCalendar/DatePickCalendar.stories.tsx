import { DatePickCalendar } from "./index";

import type { Meta, StoryObj } from "@storybook/react";

const meta: Meta<typeof DatePickCalendar> = {
  title: "shared/DatePicker/DatePickCalendar",
  component: DatePickCalendar,
};

export default meta;
type Story = StoryObj<typeof DatePickCalendar>;

export const Normal: Story = {
  args: {
    selectedDate: new Date("2023/3/10"),
  },
};
