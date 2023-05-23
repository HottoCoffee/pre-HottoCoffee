import { client } from "../../modules/aspidaClient";
import { restGet } from "~/modules/handlerFactory";
import { MOCK_BATCH_ID } from "./batch";
import { MOCK_BATCH_HISTORY_ID } from "./batchHistory";
import { add } from "date-fns";
import { components } from "~/swagger/schema/schemas/history";

/**
 * @package
 */
export const mockLength = 10;

/**
 * @package
 */
export const getHistoryMockList = (startDate: Date): Array<components["schemas"]["History"]> => {
  const list: Array<components["schemas"]["History"]> = [...new Array(mockLength)].map((_, i) => {
    return {
      history_id: MOCK_BATCH_HISTORY_ID,
      batch_id: MOCK_BATCH_ID,
      batch_name: `Batch A`,
      start_datetime: add(startDate, { days: i }).toISOString(),
      status: "success",
      finish_datetime: add(startDate, { days: i, hours: 1 }).toISOString(),
    };
  });

  return list;
};

export const mockGetBatchListForCalendar = restGet(client.api.calendar, (req, res, context) => {
  const startDate = new URLSearchParams(req.url.search).get("start_date") ?? "";

  return res(context.json(getHistoryMockList(new Date(startDate))));
});

export const calendarHandlers = [mockGetBatchListForCalendar];
