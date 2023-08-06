import { Meta, StoryObj } from "@storybook/react";
import { SignInUpForm } from ".";

const meta: Meta<typeof SignInUpForm> = {
  component: SignInUpForm,
};

export default meta;
type Story = StoryObj<typeof SignInUpForm>;

export const Default: Story = {
  args: {
    type: "sign-in",
  },
};
