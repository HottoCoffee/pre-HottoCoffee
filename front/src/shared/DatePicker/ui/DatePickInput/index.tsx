import styles from "./DatePicker.modules.scss";
import { BsCalendarEvent } from "react-icons/bs";
import * as Popover from "@radix-ui/react-popover";
import { DatePickCalendar } from "../DatePickCalendar";
import { format } from "date-fns";
import { forwardRef, Ref } from "react";

interface Props {
  selectedDate?: Date;
  onSelectDate: (date: Date) => void;
  dataTestId?: string;
}

export const DatePickInput = forwardRef(function DatePickInput(
  props: Props,
  ref: Ref<HTMLInputElement>,
) {
  const { selectedDate, onSelectDate, dataTestId } = props;

  return (
    <div className={styles.container} role="button">
      <BsCalendarEvent className={styles.icon} />
      <Popover.Root>
        <Popover.Trigger asChild>
          <input
            type="text"
            className={styles.input}
            placeholder="2023/1/3"
            value={selectedDate ? format(selectedDate, "yyyy/MM/dd") : undefined}
            ref={ref}
            data-testid={dataTestId}
          />
        </Popover.Trigger>

        <Popover.Portal>
          <Popover.Content className={styles.popoverContent}>
            <DatePickCalendar selectedDate={selectedDate} onSelectDate={onSelectDate} />
          </Popover.Content>
        </Popover.Portal>
      </Popover.Root>
    </div>
  );
});
