import { expect, jest } from "@storybook/jest";
import { EditBatchForm } from "./index";

import type { Meta, StoryObj } from "@storybook/react";
import { queryClientDecorator } from "~/modules/queryClientDecorator";
import { failedToCreateNewBatch, successToCreateNewBatch } from "~/msw/api/batch";
import { toasterDecorator } from "~/modules/toasterDecorator";
import { within, userEvent } from "@storybook/testing-library";

const meta: Meta<typeof EditBatchForm> = {
  title: "feature/BatchForm/EditBatchForm",
  component: EditBatchForm,
  decorators: [queryClientDecorator, toasterDecorator],
};

export default meta;
type Story = StoryObj<typeof EditBatchForm>;

export const SuccessToCreate: Story = {
  args: {
    onSuccess: jest.fn(),
    initialBatch: {
      id: 1,
      batch_name: "BatchA",
      server_name: "ServerB",
      cron_setting: "30 * * *",
      initial_date: "2021-01-01T00:00:00",
      time_limit: 20,
    },
  },
  parameters: {
    msw: {
      handlers: [successToCreateNewBatch()],
    },
  },
};

export const FailedToCreate: Story = {
  args: {
    onSuccess: jest.fn(),
    initialBatch: {
      id: 1,
      batch_name: "BatchA",
      server_name: "ServerB",
      cron_setting: "30 * * *",
      initial_date: "2021-01-01T00:00:00",
      time_limit: 20,
    },
  },
  parameters: {
    msw: {
      handlers: [failedToCreateNewBatch()],
    },
  },
  play: async ({ canvasElement, args }) => {
    const canvas = within(canvasElement);

    const submissionButton = canvas.getByTestId("submit");
    await userEvent.click(submissionButton);

    await expect(args.onSuccess).not.toBeCalled();
  },
};
