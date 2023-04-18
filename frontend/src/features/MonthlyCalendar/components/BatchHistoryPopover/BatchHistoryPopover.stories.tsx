import { BatchHistoryPopover } from "./index";

import type { Meta, StoryObj } from "@storybook/react";
import { queryClientDecorator } from "~/modules/queryClientDecorator";
import { MOCK_BATCH_ID } from "~/msw/api/batch";
import {
  longLoadingGetBatchHistory,
  MOCK_BATCH_HISTORY_ID,
  returnErrorGetBatchHistory,
  successGetBatchHistory,
} from "~/msw/api/batchHistory";

const meta: Meta<typeof BatchHistoryPopover> = {
  title: "feature/MonthlyCalendar/BatchHistoryPopover",
  component: BatchHistoryPopover,
  decorators: [queryClientDecorator],
};

export default meta;
type Story = StoryObj<typeof BatchHistoryPopover>;

export const Loading: Story = {
  args: {
    batchId: MOCK_BATCH_ID,
    historyId: MOCK_BATCH_HISTORY_ID,
  },
  parameters: {
    msw: {
      handlers: [longLoadingGetBatchHistory()],
    },
  },
};

export const SuccessFetching: Story = {
  args: {
    batchId: MOCK_BATCH_ID,
    historyId: MOCK_BATCH_HISTORY_ID,
  },
  parameters: {
    msw: {
      handlers: [successGetBatchHistory()],
    },
  },
};

export const ApiError: Story = {
  args: {
    batchId: MOCK_BATCH_ID,
    historyId: MOCK_BATCH_HISTORY_ID,
  },
  parameters: {
    msw: {
      handlers: [returnErrorGetBatchHistory()],
    },
  },
};
