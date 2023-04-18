import { expect } from "@storybook/jest";
import { DatePickCalendar } from "./index";
import * as Popover from "@radix-ui/react-popover";
import { action } from "@storybook/addon-actions";

import type { Meta, StoryObj } from "@storybook/react";
import { userEvent, within } from "@storybook/testing-library";

const meta = {
  title: "shared/DatePicker/DatePickCalendar",
  component: DatePickCalendar,
  decorators: [
    (story) => {
      return <Popover.Root>{story()}</Popover.Root>;
    },
  ],
} satisfies Meta<typeof DatePickCalendar>;

export default meta;
type Story = StoryObj<typeof DatePickCalendar>;

export const Normal: Story = {
  args: {
    selectedDate: new Date("2023/3/10"),
    onSelectDate: action("select date"),
  },
  play: async (context) => {
    const canvas = within(context.canvasElement);

    const backToPreviousMonthButton = canvas.getByTestId("back-previous-month");
    await userEvent.click(backToPreviousMonthButton);
    await expect(canvas.getByTestId("date-input-title").innerText).toBe("2023/02");

    const goToNextMonthButton = canvas.getByTestId("go-next-month");
    await userEvent.click(goToNextMonthButton);
    await expect(canvas.getByTestId("date-input-title").innerText).toBe("2023/03");

    await userEvent.click(goToNextMonthButton);
    await expect(canvas.getByTestId("date-input-title").innerText).toBe("2023/04");
  },
};
