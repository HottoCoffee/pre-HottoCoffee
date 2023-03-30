import { TextInput } from "./index";

import type { Meta, StoryObj } from "@storybook/react";
import { queryClientDecorator } from "~/modules/queryClientDecorator";

const meta: Meta<typeof TextInput> = {
  title: "shared/TextInput",
  component: TextInput,
  decorators: [queryClientDecorator],
};

export default meta;
type Story = StoryObj<typeof TextInput>;

export const Empty: Story = {
  args: {},
};

export const WithPlaceHolder: Story = {
  args: {
    placeholder: "Placeholder",
  },
};
