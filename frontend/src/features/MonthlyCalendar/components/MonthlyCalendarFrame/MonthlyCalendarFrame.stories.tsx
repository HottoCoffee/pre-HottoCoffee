import { MonthlyCalendarFrame } from "./index";

import type { Meta, StoryObj } from "@storybook/react";

const meta: Meta<typeof MonthlyCalendarFrame> = {
  title: "feature/MonthlyCalendar/MonthlyCalendarFrame",
  component: MonthlyCalendarFrame,
  argTypes: {
    date: {
      control: {
        type: "date",
      },
    },
  },
};

export default meta;
type Story = StoryObj<typeof MonthlyCalendarFrame>;

export const SimpleCalendar: Story = {
  args: {
    date: new Date("2023/3/1"),
    children: (date: Date) => {
      return (
        <div>
          <p>{date.getDate()}</p>
        </div>
      );
    },
  },
};
