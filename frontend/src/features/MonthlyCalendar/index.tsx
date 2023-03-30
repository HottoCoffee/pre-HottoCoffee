import { useAspidaQuery } from "@aspida/react-query";
import { MonthlyCalendarFrame } from "./components/MonthlyCalendarFrame";
import { client } from "~/modules/aspidaClient";
import { getDateListInMonth } from "./utils/dateCalculationHelper";
import { isSameDay } from "date-fns";
import { ReactNode } from "react";

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
    <div>
      <MonthlyCalendarFrame date={date}>
        {(date) => {
          const targetBatchList = data?.filter(({ start_datetime }) => {
            return isSameDay(new Date(start_datetime), date);
          });
          return <div>{targetBatchList?.at(0)?.start_datetime}</div>;
        }}
      </MonthlyCalendarFrame>
    </div>
  );
};
