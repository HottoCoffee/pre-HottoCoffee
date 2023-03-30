import { BatchDetailPopover } from "./index";

import type { Meta, StoryObj } from "@storybook/react";
import { queryClientDecorator } from "~/modules/queryClientDecorator";
import {
  longLoadingGetBatchByBatchId,
  MOCK_BATCH_ID,
  returnErrorGetBatchByBatchId,
  successGetBatchByBatchId,
} from "~/msw/api/batch";

const meta: Meta<typeof BatchDetailPopover> = {
  title: "feature/MonthlyCalendar/BatchDetailPopover",
  component: BatchDetailPopover,
  decorators: [queryClientDecorator],
};

export default meta;
type Story = StoryObj<typeof BatchDetailPopover>;

export const Loading: Story = {
  args: {
    batchId: MOCK_BATCH_ID,
  },
  parameters: {
    msw: {
      handlers: [longLoadingGetBatchByBatchId()],
    },
  },
};

export const SuccessFetching: Story = {
  args: {
    batchId: MOCK_BATCH_ID,
  },
  parameters: {
    msw: {
      handlers: [successGetBatchByBatchId()],
    },
  },
};

export const ApiError: Story = {
  args: {
    batchId: MOCK_BATCH_ID,
  },
  parameters: {
    msw: {
      handlers: [returnErrorGetBatchByBatchId()],
    },
  },
};
