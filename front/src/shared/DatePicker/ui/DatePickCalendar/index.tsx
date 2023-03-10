import { add, format, isSameDay, isSameMonth } from "date-fns";
import { useState } from "react";
import { getDateListInMonth } from "~/features/MonthlyCalendar/utils/dateCalculationHelper";

import styles from "./DatePickCalendar.modules.scss";
import * as AccessibleIcon from "@radix-ui/react-accessible-icon";
import { MdOutlineArrowBackIos, MdOutlineArrowForwardIos } from "react-icons/md";

interface Props {
  selectedDate?: Date;
  onSelectDate: (date: Date) => void;
}

export const DatePickCalendar = (props: Props) => {
  const { selectedDate = new Date(), onSelectDate } = props;
  const [displayedMonth, setDisplayedMonth] = useState(selectedDate);
  const dateList = getDateListInMonth(displayedMonth);

  const onSelectDay = (date: Date) => {
    setDisplayedMonth(date);
    onSelectDate(date);
  };

  const goPreviousMonth = () => {
    setDisplayedMonth(add(displayedMonth, { months: -1 }));
  };

  const goNextMonth = () => {
    setDisplayedMonth(add(displayedMonth, { months: 1 }));
  };

  return (
    <div className={styles.wrapper}>
      <div className={styles.displayedMonthLabel}>
        <AccessibleIcon.Root label="BackToPreviousMonth">
          <button className={styles.arrowButton} onClick={goPreviousMonth}>
            <MdOutlineArrowBackIos />
          </button>
        </AccessibleIcon.Root>

        <p>{format(displayedMonth, "yyyy/MM")}</p>

        <AccessibleIcon.Root label="BackToPreviousMonth">
          <button className={styles.arrowButton} onClick={goNextMonth}>
            <MdOutlineArrowForwardIos />
          </button>
        </AccessibleIcon.Root>
      </div>

      <div className={styles.container}>
        {dateList.map((week, weekIndex) => {
          return (
            <div key={weekIndex} className={styles.week}>
              {week.map((date, dayIndex) => {
                return (
                  <button
                    key={weekIndex * 7 + dayIndex}
                    className={styles.day}
                    data-isthismonth={isSameMonth(date, displayedMonth)}
                    data-isselected={isSameDay(date, selectedDate)}
                    onClick={() => {
                      onSelectDay(date);
                    }}
                  >
                    {format(date, "d")}
                  </button>
                );
              })}
            </div>
          );
        })}
      </div>
    </div>
  );
};
