import { calculateRenderTarget } from "./functions";

describe("features > MonthlyCalendar > components > MonthlyCalendarDayCell > functions", () => {
  describe("calculateRenderTarget", () => {
    test("should return { hasMore: false } when batchListSize is 0", () => {
      expect(calculateRenderTarget(0, 100)).toEqual({ hasMore: false });
    });
    test("should return { hasMore: false } when batchListSize is 1", () => {
      expect(calculateRenderTarget(1, 100)).toEqual({ hasMore: false });
    });
    test("should return { hasMore: false } when batchListSize is 2", () => {
      expect(calculateRenderTarget(2, 100)).toEqual({ hasMore: false });
    });
    test("should return { hasMore: true, showCount: 0, restCount: 2 } when batchListSize is 2 and allocatedAreaHeight is 50", () => {
      expect(calculateRenderTarget(2, 50)).toEqual({ hasMore: true, showCount: 0, restCount: 2 });
    });
    test("should return { hasMore: true, showCount: 0, restCount: 3 } when batchListSize is 3 and allocatedAreaHeight is 50", () => {
      expect(calculateRenderTarget(3, 50)).toEqual({ hasMore: true, showCount: 0, restCount: 3 });
    });
    test("should return { hasMore: true, showCount: 1, restCount: 2 } when batchListSize is 3 and allocatedAreaHeight is 80", () => {
      expect(calculateRenderTarget(3, 80)).toEqual({ hasMore: true, showCount: 1, restCount: 2 });
    });
  });
});
