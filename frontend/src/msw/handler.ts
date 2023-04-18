import { batchHandlers } from "./api/batch";
import { batchHistoryHandlers } from "./api/batchHistory";
import { calendarHandlers } from "./api/calendar";

export const handlers = [...calendarHandlers, ...batchHandlers, ...batchHistoryHandlers];
