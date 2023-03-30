import { useAspidaQuery } from "@aspida/react-query";
import { MonthlyCalendarFrame } from "./components/MonthlyCalendarFrame";
import { client } from "~/modules/aspidaClient";
import { getDateListInMonth } from "./utils/dateCalculationHelper";
import { isSameDay } from "date-fns";
import { BatchStatusButton } from "./components/BatchStatusButton";
import styles from "./index.module.scss";
import classNames from "classnames/bind";
import { MonthlyCalendarDayCell } from "./components/MonthlyCalendarDayCell";

const clx = classNames.bind(styles);

interface Props {
  date: Date;
}

export const MonthlyCalendar = (props: Props) => {
  const { date } = props;

  const dateList = getDateListInMonth(date);
  const startDate = dateList[0][0];
  const endDate = dateList.at(-1)?.at(-1);

  const { data, error } = useAspidaQuery(client.api.calendar, {
    query: {
      start_date: startDate.toISOString() ?? "",
      end_date: endDate?.toISOString() ?? "",
    },
  });

  return (
    <div className={styles.container}>
      <MonthlyCalendarFrame date={date}>
        {(date) => {
          const targetBatchList = data?.filter(({ start_datetime }) => {
            return isSameDay(new Date(start_datetime), date);
          });
          if (!targetBatchList) {
            return null;
          }

          return <MonthlyCalendarDayCell batchHistoryList={targetBatchList} date={date} />;
        }}
      </MonthlyCalendarFrame>
    </div>
  );
};
