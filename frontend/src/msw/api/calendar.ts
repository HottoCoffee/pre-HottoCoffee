import { client } from "../../modules/aspidaClient";
import { restGet } from "~/modules/handlerFactory";

export const mockGetBatchListForCalendar = restGet(client.api.batch, (req, res, context) => {
  return res(
    context.json([
      {
        id: 1,
        batch_name: "Batch1",
        server_name: "Server1",
        initial_date: "2023-03-03T09:27:03.529Z",
        time_limit: 30,
        cron_setting: "30 * * *",
      },
    ]),
  );
});

export const calendarHandlers = [mockGetBatchListForCalendar];
