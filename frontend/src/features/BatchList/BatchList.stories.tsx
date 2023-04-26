import { Meta, StoryObj } from "@storybook/react";
import { BatchList } from ".";
import { successGetBatchList, successToCreateNewBatch } from "~/msw/api/batch";
import { queryClientDecorator } from "~/modules/queryClientDecorator";
import { toasterDecorator } from "~/modules/toasterDecorator";

const meta: Meta<typeof BatchList> = {
  title: "feature/BatchList",
  component: BatchList,
  decorators: [queryClientDecorator, toasterDecorator],
};

export default meta;
type Story = StoryObj<typeof BatchList>;

export const Default: Story = {
  parameters: {
    msw: {
      handlers: [successToCreateNewBatch(), successGetBatchList()],
    },
  },
};
