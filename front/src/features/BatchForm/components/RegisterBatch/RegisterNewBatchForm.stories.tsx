import { RegisterNewBatchForm } from "./index";

import type { Meta, StoryObj } from "@storybook/react";
import { queryClientDecorator } from "~/modules/queryClientDecorator";

const meta: Meta<typeof RegisterNewBatchForm> = {
  title: "feature/BatchForm/RegisterNewBatchForm",
  component: RegisterNewBatchForm,
  decorators: [queryClientDecorator],
};

export default meta;
type Story = StoryObj<typeof RegisterNewBatchForm>;

export const Loading: Story = { args: {} };
