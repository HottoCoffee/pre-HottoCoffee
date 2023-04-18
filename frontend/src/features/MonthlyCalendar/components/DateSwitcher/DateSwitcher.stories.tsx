import { DateSwitcher } from "./index";

import type { Meta, StoryObj } from "@storybook/react";
import { queryClientDecorator } from "~/modules/queryClientDecorator";

const meta = {
  title: "feature/MonthlyCalendar/DateSwitcher",
  component: DateSwitcher,
  decorators: [queryClientDecorator],
} satisfies Meta<typeof DateSwitcher>;

export default meta;
type Story = StoryObj<typeof DateSwitcher>;

export const Loading: Story = {
  args: {
    type: "month",
    date: new Date(),
  },
};
