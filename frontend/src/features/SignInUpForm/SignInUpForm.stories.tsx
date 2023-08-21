import { Meta, StoryObj } from "@storybook/react";
import { SignInUpForm } from ".";
import { queryClientDecorator } from "~/modules/queryClientDecorator";
import { toasterDecorator } from "~/modules/toasterDecorator";

const meta: Meta<typeof SignInUpForm> = {
  component: SignInUpForm,
  decorators: [queryClientDecorator, toasterDecorator],
};

export default meta;
type Story = StoryObj<typeof SignInUpForm>;

export const Default: Story = {
  args: {
    type: "sign-in",
  },
};
