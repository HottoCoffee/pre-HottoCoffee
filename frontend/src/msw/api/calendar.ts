import { client } from "../../modules/aspidaClient";
import { restGet } from "~/modules/handlerFactory";

export const mockGetBatchListForCalendar = restGet(client.api.calendar, (req, res, context) => {
  return res(
    context.json([
      {
        history_id: 1,
        batch_id: 1,
        batch_name: "Batch 1",
        start_datetime: "2023-02-04T00:00:00.000Z",
        status: "success",
      },
      {
        history_id: 1,
        batch_id: 1,
        batch_name: "LongLongBatchName",
        start_datetime: "2023-02-04T00:00:00.000Z",
        status: "success",
      },
      {
        history_id: 1,
        batch_id: 1,
        batch_name: "Batch 1",
        start_datetime: "2023-02-04T00:00:00.000Z",
        status: "success",
      },
    ]),
  );
});

export const calendarHandlers = [mockGetBatchListForCalendar];
