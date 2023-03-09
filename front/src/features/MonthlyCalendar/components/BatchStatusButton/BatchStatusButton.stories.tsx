import { BatchStatusButton } from "./index";

import type { Meta, StoryObj } from "@storybook/react";
import { MOCK_BATCH_ID, successGetBatchByBatchId } from "~/msw/api/batch";
import { queryClientDecorator } from "~/modules/queryClientDecorator";
import { MOCK_BATCH_HISTORY_ID, successGetBatchHistory } from "~/msw/api/batchHistory";

const meta: Meta<typeof BatchStatusButton> = {
  title: "feature/MonthlyCalendar/BatchStatusButton",
  component: BatchStatusButton,
  decorators: [queryClientDecorator],
};

export default meta;
type Story = StoryObj<typeof BatchStatusButton>;

export const BeforeStarted: Story = {
  args: {
    batchName: "Batch A",
    status: "before_started",
    batchId: MOCK_BATCH_ID,
  },
  parameters: {
    msw: {
      handlers: [successGetBatchByBatchId()],
    },
  },
};

export const InProgress: Story = {
  args: {
    batchName: "Batch A",
    status: "in_progress",
    batchId: MOCK_BATCH_ID,
    historyId: MOCK_BATCH_HISTORY_ID,
  },
  parameters: {
    msw: {
      handlers: [successGetBatchHistory("in_progress")],
    },
  },
};

export const Success: Story = {
  args: {
    batchName: "Batch A",
    status: "success",
    batchId: MOCK_BATCH_ID,
    historyId: MOCK_BATCH_HISTORY_ID,
  },
  parameters: {
    msw: {
      handlers: [successGetBatchHistory("success")],
    },
  },
};

export const Failed: Story = {
  args: {
    batchName: "Batch A",
    status: "failed",
    batchId: MOCK_BATCH_ID,
    historyId: MOCK_BATCH_HISTORY_ID,
  },
  parameters: {
    msw: {
      handlers: [successGetBatchHistory("failed")],
    },
  },
};
