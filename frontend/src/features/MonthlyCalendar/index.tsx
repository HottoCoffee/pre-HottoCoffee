import { useAspidaQuery } from "@aspida/react-query";
import { MonthlyCalendarFrame } from "./components/MonthlyCalendarFrame";
import { client } from "~/modules/aspidaClient";
import { getDateListInMonth } from "./utils/dateCalculationHelper";
import { isSameDay } from "date-fns";
import styles from "./index.module.scss";
import { MonthlyCalendarDayCell } from "./components/MonthlyCalendarDayCell";
import { Toaster } from "~/shared/Toaster/ui";
import * as Toast from "@radix-ui/react-toast";

interface Props {
  date: Date;
}

export const MonthlyCalendar = (props: Props) => {
  const { date } = props;

  const dateList = getDateListInMonth(date);
  const startDate = dateList[0][0];
  const endDate = dateList.at(-1)?.at(-1);

  console.log(startDate.toISOString() ?? "");

  const { data, error } = useAspidaQuery(client.api.calendar, {
    query: {
      start_date: startDate.toISOString() ?? "",
      end_date: endDate?.toISOString() ?? "",
    },
  });

  // FIXME: Should add error handling
  if (error) {
    return <p>Unknown error occurred. {JSON.stringify(error)}</p>
  }

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

      {error && (
        <>
          <Toaster
            type={"success"}
            description={<p>{error}</p>}
            title="Error on api."
            open={Boolean(error)}
          />
          <Toast.Viewport />
        </>
      )}
    </div>
  );
};
