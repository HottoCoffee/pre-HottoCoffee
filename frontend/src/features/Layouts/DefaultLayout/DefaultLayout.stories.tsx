import { Meta, StoryObj } from "@storybook/react";
import { DefaultLayout } from ".";
import { queryClientDecorator } from "~/modules/queryClientDecorator";
import { toasterDecorator } from "~/modules/toasterDecorator";

const meta: Meta<typeof DefaultLayout> = {
  title: "feature/layouts/DefaultLayout",
  component: DefaultLayout,
  decorators: [queryClientDecorator, toasterDecorator],
};

export default meta;
type Story = StoryObj<typeof DefaultLayout>;

export const Default: Story = {
  args: {
    children: <div>children</div>,
    headerChildren: <div>headerChildren</div>,
  },
};
