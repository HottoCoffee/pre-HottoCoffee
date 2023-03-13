import { DatePickCalendar } from "./index";
import * as Popover from "@radix-ui/react-popover";
import { action } from "@storybook/addon-actions";

import type { Meta, StoryObj } from "@storybook/react";
import { userEvent, within } from "@storybook/testing-library";

const meta: Meta<typeof DatePickCalendar> = {
  title: "shared/DatePicker/DatePickCalendar",
  component: DatePickCalendar,
  decorators: [
    (story) => {
      return <Popover.Root>{story()}</Popover.Root>;
    },
  ],
};

export default meta;
type Story = StoryObj<typeof DatePickCalendar>;

export const Normal: Story = {
  args: {
    selectedDate: new Date("2023/3/10"),
    onSelectDate: action("select date"),
  },
  play: async ({ canvasElement }) => {
    const canvas = within(canvasElement);

    const backToPreviousMonthButton = canvas.getByTestId("back-previous-month");
    await userEvent.click(backToPreviousMonthButton);
  },
};
