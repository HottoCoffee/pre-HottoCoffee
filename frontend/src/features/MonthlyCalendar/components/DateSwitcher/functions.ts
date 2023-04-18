import { match } from "ts-pattern";
import { Props } from "./index";
import { add } from "date-fns";

export const dateSwitcherOnChangeMethod = (props: Props) => {
  const { date, type } = props;

  return match(type)
    .with("month", () => {
      return {
        goNext: () => {
          return add(date, { months: 1 });
        },
        backPrevious: () => {
          return add(date, { months: -1 });
        },
      };
    })
    .with("week", () => {
      return {
        goNext: () => {
          return add(date, { weeks: 1 });
        },
        backPrevious: () => {
          return add(date, { weeks: -1 });
        },
      };
    })
    .exhaustive();
};
