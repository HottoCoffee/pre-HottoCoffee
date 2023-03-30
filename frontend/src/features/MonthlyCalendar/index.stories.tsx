import { mockGetBatchListForCalendar } from "~/msw/api/calendar";
import { MonthlyCalendar } from "./index";

import type { Meta, StoryObj } from "@storybook/react";
import { queryClientDecorator } from "~/modules/queryClientDecorator";
import { successGetBatchByBatchId } from "~/msw/api/batch";
import { toasterDecorator } from "~/modules/toasterDecorator";
import { successGetBatchHistory } from "~/msw/api/batchHistory";

const meta: Meta<typeof MonthlyCalendar> = {
  title: "feature/MonthlyCalendar",
  component: MonthlyCalendar,
  decorators: [queryClientDecorator, toasterDecorator],
};

export default meta;
type Story = StoryObj<typeof MonthlyCalendar>;

export const Default: Story = {
  args: {
    date: new Date("2023/2/1"),
  },
  parameters: {
    msw: {
      handlers: [mockGetBatchListForCalendar, successGetBatchByBatchId(), successGetBatchHistory()],
    },
  },
};
