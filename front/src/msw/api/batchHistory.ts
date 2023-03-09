import { rest } from "msw";
import { client } from "~/modules/aspidaClient";
import { restGet } from "~/modules/handlerFactory";
import { sleep } from "~/modules/sleep";
import { HottoCoffee } from "~/types";
import { MOCK_BATCH_ID } from "./batch";

export const MOCK_BATCH_HISTORY_ID = 1;

export const successGetBatchHistory = (status: HottoCoffee.BatchStatus = "success") => {
  return restGet(
    client.api.batch._batch_id(MOCK_BATCH_ID).history._history_id(MOCK_BATCH_HISTORY_ID),
    async (_, res, context) => {
      return res(
        context.json({
          batch_id: MOCK_BATCH_ID,
          history_id: MOCK_BATCH_HISTORY_ID,
          batch_name: "BatchA",
          start_datetime: new Date("2023/3/14").toISOString(),
          status: status,
        }),
      );
    },
  );
};

export const longLoadingGetBatchHistory = () => {
  return restGet(
    client.api.batch._batch_id(MOCK_BATCH_ID).history._history_id(MOCK_BATCH_HISTORY_ID),
    async () => {
      await sleep(1000);
    },
  );
};

export const returnErrorGetBatchHistory = () => {
  return rest.get(
    client.api.batch._batch_id(MOCK_BATCH_ID).history._history_id(MOCK_BATCH_HISTORY_ID).$path(),
    async (_, res, context) => {
      return res(
        context.status(500),
        context.json({
          status: 500,
          message: "Unknown Error",
        }),
      );
    },
  );
};

export const batchHandlers = [successGetBatchHistory()];
