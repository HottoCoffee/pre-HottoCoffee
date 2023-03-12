/**
 * This file was auto-generated by openapi-typescript.
 * Do not make direct changes to the file.
 */


export type paths = Record<string, never>;

export type webhooks = Record<string, never>;

export interface components {
  schemas: {
    History: {
      /**
       * Format: int64 
       * @description history_id will be null when the batch execution time is future.
       */
      history_id?: number;
      /** Format: int64 */
      batch_id: number;
      batch_name: string;
      /** Format: ISO8601 */
      start_datetime: string;
      /** @enum {string} */
      status: "before_started" | "in_progress" | "success" | "failed";
    };
  };
  responses: never;
  parameters: never;
  requestBodies: never;
  headers: never;
  pathItems: never;
}

export type external = Record<string, never>;

export type operations = Record<string, never>;
