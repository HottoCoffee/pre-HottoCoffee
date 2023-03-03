import { MonthlyCalendarFrame } from "./index";

import type { Meta, StoryObj } from "@storybook/react";

const meta: Meta<typeof MonthlyCalendarFrame> = {
  title: "feature/MonthlyCalendar/MonthlyCalendarFrame",
  component: MonthlyCalendarFrame,
};

export default meta;
type Story = StoryObj<typeof MonthlyCalendarFrame>;

export const HasNoChild: Story = {
  args: {
    date: new Date("2023/3/1"),
    childBoxComponent: (date: Date) => {
      return (
        <div>
          <p>{date.getDate()}</p>
        </div>
      );
    },
  },
};
