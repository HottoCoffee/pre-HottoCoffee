import styles from "./MonthlyCalendarFrame.modules.scss";
import { getDateListInMonth } from "../../utils/dateCalculationHelper";

interface Props {
  date: Date;
  childBoxComponent: (date: Date) => React.ReactNode;
}

export const MonthlyCalendarFrame = (props: Props) => {
  const { date, childBoxComponent } = props;

  const dateList = getDateListInMonth(date);

  return (
    <div className={styles.container}>
      {dateList.map((week, weekIndex) => {
        return (
          <div key={weekIndex} className={styles.week}>
            {week.map((date, dayIndex) => {
              return (
                <div key={weekIndex * 7 + dayIndex} className={styles.day}>
                  {childBoxComponent(date)}
                </div>
              );
            })}
          </div>
        );
      })}
    </div>
  );
};
