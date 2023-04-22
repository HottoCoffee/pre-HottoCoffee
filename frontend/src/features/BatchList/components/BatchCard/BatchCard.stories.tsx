import { Meta, StoryObj } from "@storybook/react";
import { BatchCard } from ".";

const meta: Meta<typeof BatchCard> = {
  title: "feature/BatchList/BatchCard",
  component: BatchCard,
};

export default meta;
type Story = StoryObj<typeof BatchCard>;

export const SuccessToFetch: Story = {
  args: {
    batch: {
      id: 1,
      batch_name: "エキスポート",
      server_name: "初号機",
      cron_setting: "0 0 * * *",
      initial_date: "2023-01-01",
      time_limit: 60,
    },
  },
};
