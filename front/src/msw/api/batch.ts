import { rest } from "msw";
import { client } from "~/modules/aspidaClient";
import { restGet } from "~/modules/handlerFactory";
import { sleep } from "~/modules/sleep";

export const MOCK_BATCH_ID = 1;

export const successGetBatchByBatchId = () => {
  return restGet(client.api.batch._batch_id(MOCK_BATCH_ID), async (_, res, context) => {
    return res(
      context.json({
        id: MOCK_BATCH_ID,
        batch_name: "Batch1",
        server_name: "Server1",
        initial_date: "2023-03-03T09:27:03.529Z",
        time_limit: 30,
        cron_setting: "30 * * *",
      }),
    );
  });
};

export const longLoadingGetBatchByBatchId = () => {
  return restGet(client.api.batch._batch_id(MOCK_BATCH_ID), async () => {
    await sleep(1000);
  });
};

export const returnErrorGetBatchByBatchId = () => {
  return rest.get(client.api.batch._batch_id(MOCK_BATCH_ID).$path(), async (_, res, context) => {
    return res(
      context.status(500),
      context.json({
        status: 500,
        message: "Unknown Error",
      }),
    );
  });
};

export const batchHandlers = [successGetBatchByBatchId()];
