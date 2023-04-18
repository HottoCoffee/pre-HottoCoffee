import { RegisterNewBatchWithModal } from "./index";

import type { Meta, StoryObj } from "@storybook/react";
import { queryClientDecorator } from "~/modules/queryClientDecorator";
import { successToCreateNewBatch } from "~/msw/api/batch";
import { toasterDecorator } from "~/modules/toasterDecorator";
import { within, userEvent } from "@storybook/testing-library";

const meta: Meta<typeof RegisterNewBatchWithModal> = {
  title: "feature/BatchForm/RegisterNewBatchWithModal",
  component: RegisterNewBatchWithModal,
  decorators: [queryClientDecorator, toasterDecorator],
};

export default meta;
type Story = StoryObj<typeof RegisterNewBatchWithModal>;

export const SuccessToCreate: Story = {
  args: {
    children: <button>Open</button>,
  },
  parameters: {
    msw: {
      handlers: [successToCreateNewBatch()],
    },
  },
  play: async ({ canvasElement, args }) => {
    const canvas = within(canvasElement);

    const openButton = canvas.getByText("Open");
    await userEvent.click(openButton);
  },
};
