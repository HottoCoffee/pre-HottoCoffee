import { composeStories } from "@storybook/react";
import * as stories from "./DatePickCalendar.stories";

const { Normal } = composeStories(stories);

describe("Should show previous month", () => {
  test("should go previous month", () => {});
});
