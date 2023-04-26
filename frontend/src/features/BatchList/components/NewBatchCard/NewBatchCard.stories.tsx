import { Meta, StoryObj } from "@storybook/react";
import { NewBatchCard } from ".";
import { successToCreateNewBatch } from "~/msw/api/batch";
import { queryClientDecorator } from "~/modules/queryClientDecorator";
import { toasterDecorator } from "~/modules/toasterDecorator";

const meta: Meta<typeof NewBatchCard> = {
  title: "feature/BatchList/NewBatchCard",
  component: NewBatchCard,
  decorators: [queryClientDecorator, toasterDecorator],
};

export default meta;
type Story = StoryObj<typeof NewBatchCard>;

export const Default: Story = {
  parameters: {
    msw: {
      handlers: [successToCreateNewBatch()],
    },
  },
};
