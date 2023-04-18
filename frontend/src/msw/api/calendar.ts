import { client } from "../../modules/aspidaClient";
import { restGet } from "~/modules/handlerFactory";
import { MOCK_BATCH_ID } from "./batch";
import { MOCK_BATCH_HISTORY_ID } from "./batchHistory";

export const mockGetBatchListForCalendar = restGet(client.api.calendar, (req, res, context) => {
  return res(
    context.json([
      {
        history_id: MOCK_BATCH_HISTORY_ID,
        batch_id: MOCK_BATCH_ID,
        batch_name: "Batch 1",
        start_datetime: "2023-02-04T00:00:00.000Z",
        status: "success",
      },
      {
        history_id: MOCK_BATCH_HISTORY_ID,
        batch_id: MOCK_BATCH_ID,
        batch_name: "LongLongBatchName",
        start_datetime: "2023-02-04T00:00:00.000Z",
        status: "success",
      },
      {
        history_id: MOCK_BATCH_HISTORY_ID,
        batch_id: MOCK_BATCH_ID,
        batch_name: "Batch 1",
        start_datetime: "2023-02-04T00:00:00.000Z",
        status: "success",
      },
    ]),
  );
});

export const calendarHandlers = [mockGetBatchListForCalendar];
