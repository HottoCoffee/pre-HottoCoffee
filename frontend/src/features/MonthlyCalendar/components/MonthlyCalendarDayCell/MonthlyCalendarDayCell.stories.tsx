import { MonthlyCalendarDayCell } from "./index";

import type { Meta, StoryObj } from "@storybook/react";

const meta: Meta<typeof MonthlyCalendarDayCell> = {
  title: "feature/MonthlyCalendar/MonthlyCalendarDayCell",
  component: MonthlyCalendarDayCell,
};

export default meta;
type Story = StoryObj<typeof MonthlyCalendarDayCell>;

export const Default: Story = {
  args: {
    date: new Date("2023/2/1"),
    batchHistoryList: [
      {
        batch_id: 1,
        batch_name: "batch_name_1",
        status: "before_started",
        start_datetime: new Date("2023/2/1 10:00").toISOString(),
        finish_datetime: new Date("2023/2/1 12:00").toISOString(),
      },
      {
        batch_id: 1,
        batch_name: "batch_name_2",
        status: "before_started",
        start_datetime: new Date("2023/2/1 10:00").toISOString(),
        finish_datetime: new Date("2023/2/1 12:00").toISOString(),
      },
    ],
  },
};

export const SmallAreaAssigned: Story = {
  args: {
    date: new Date("2023/2/1"),
    batchHistoryList: [
      {
        batch_id: 1,
        batch_name: "batch_name_1",
        status: "before_started",
        start_datetime: new Date("2023/2/1 10:00").toISOString(),
        finish_datetime: new Date("2023/2/1 12:00").toISOString(),
      },
      {
        batch_id: 1,
        batch_name: "batch_name_2",
        status: "before_started",
        start_datetime: new Date("2023/2/1 10:00").toISOString(),
        finish_datetime: new Date("2023/2/1 12:00").toISOString(),
      },
      {
        batch_id: 1,
        batch_name: "batch_name_2",
        status: "before_started",
        start_datetime: new Date("2023/2/1 10:00").toISOString(),
        finish_datetime: new Date("2023/2/1 12:00").toISOString(),
      },
      {
        batch_id: 1,
        batch_name: "batch_name_2",
        status: "before_started",
        start_datetime: new Date("2023/2/1 10:00").toISOString(),
        finish_datetime: new Date("2023/2/1 12:00").toISOString(),
      },
    ],
  },
  render: (args) => {
    return (
      <div style={{ width: "100px", height: "100px" }}>
        <MonthlyCalendarDayCell {...args} />
      </div>
    );
  },
};
