import { batchHandlers } from "./api/batch";
import { calendarHandlers } from "./api/calendar";

export const handlers = [...calendarHandlers, ...batchHandlers];
