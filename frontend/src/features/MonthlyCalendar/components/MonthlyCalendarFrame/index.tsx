import styles from "./MonthlyCalendarFrame.module.scss";
import { getDateListInMonth } from "../../utils/dateCalculationHelper";

interface Props {
  date: Date;
  children: (date: Date) => React.ReactNode;
}

export const MonthlyCalendarFrame = (props: Props) => {
  const { date, children } = props;

  const dateList = getDateListInMonth(date);

  return (
    <div
      className={styles.container}
      style={{
        gridTemplateRows: `repeat(${dateList.length}, 1fr)`,
      }}
    >
      {dateList.map((week, weekIndex) => {
        return (
          <div key={weekIndex} className={styles.week}>
            {week.map((date, dayIndex) => {
              return (
                <div key={weekIndex * 7 + dayIndex} className={styles.day}>
                  {children(date)}
                </div>
              );
            })}
          </div>
        );
      })}
    </div>
  );
};
