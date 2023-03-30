import styles from "./DatePicker.modules.scss";
import { BsCalendarEvent } from "react-icons/bs";
import * as Popover from "@radix-ui/react-popover";
import { DatePickCalendar } from "../DatePickCalendar";
import { format } from "date-fns";
import { forwardRef, Ref } from "react";
import classNames from "classnames/bind";

const clx = classNames.bind(styles);

interface Props {
  selectedDate?: Date;
  onSelectDate: (date: Date) => void;
  dataTestId?: string;
}

export const DatePickInput = forwardRef(function DatePickInput(
  props: Props,
  ref: Ref<HTMLButtonElement>,
) {
  const { selectedDate, onSelectDate, dataTestId } = props;

  const buttonClass = clx(styles.button, {
    [styles.placeholder]: !selectedDate,
  });
  const value = selectedDate ?? new Date();

  return (
    <div className={styles.container}>
      <BsCalendarEvent className={styles.icon} />
      <Popover.Root>
        <Popover.Trigger asChild>
          <button className={buttonClass} ref={ref} data-testid={dataTestId}>
            {format(value, "yyyy/MM/dd")}
          </button>
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
