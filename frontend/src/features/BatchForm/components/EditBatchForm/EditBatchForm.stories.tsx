import { expect, jest } from "@storybook/jest";
import { EditBatchForm } from "./index";

import type { Meta, StoryObj } from "@storybook/react";
import { queryClientDecorator } from "~/modules/queryClientDecorator";
import { failedToCreateNewBatch, successToCreateNewBatch } from "~/msw/api/batch";
import { toasterDecorator } from "~/modules/toasterDecorator";
import { within, userEvent } from "@storybook/testing-library";
import { sleep } from "~/modules/sleep";

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
  play: async ({ canvasElement, args }) => {
    const canvas = within(canvasElement);

    // const batchNameInput = canvas.getAllByLabelText("Batch name")[0];
    // await userEvent.type(batchNameInput, "BatchA");

    // const serverNameInput = canvas.getAllByLabelText("Server name")[0];
    // await userEvent.type(serverNameInput, "ServerB");

    // const cronSettingInput = canvas.getAllByLabelText("Cron setting")[0];
    // await userEvent.type(cronSettingInput, "30 * * *");

    // const initialExecutionDateInput = canvas.getByTestId("initial_execution_date");
    // await userEvent.click(initialExecutionDateInput);
    // await userEvent.click(initialExecutionDateInput);

    // const timeLimitInput = canvas.getByLabelText("Time limit (min)");
    // await userEvent.type(timeLimitInput, "20");

    // const submissionButton = canvas.getByTestId("submit");
    // await userEvent.click(submissionButton);
    // await sleep(200);

    await expect(args.onSuccess).toBeCalled();
  },
};

export const FailedToCreate: Story = {
  args: {
    onSuccess: jest.fn(),
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
