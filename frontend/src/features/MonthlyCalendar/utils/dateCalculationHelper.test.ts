import {
  getDateListInMonth,
  getFirstDateOfMonthlyCalendar,
  getWeekCount,
} from "./dateCalculationHelper";

describe("dateCalculationHelper", () => {
  it("return 5 when date was provided 2023/3/1", () => {
    const date = new Date("2023/3/1");
    const count = getWeekCount(date);
    expect(count).toBe(5);
  });

  it("return 6 when date was provided 2023/7/1", () => {
    const date = new Date("2023/7/1");
    const count = getWeekCount(date);
    expect(count).toBe(6);
  });

  it("returns date list in 2023/3/1", () => {
    const date = new Date("2023/3/1");
    const dateList = getDateListInMonth(date);

    expect(dateList.length).toBe(5);
    expect(dateList[0].length).toBe(7);
  });

  it("returns 2023/2/26 as first date of monthly calendar when 2023/3/1 was provided ", () => {
    const date = new Date("2023/3/1");
    const firstDate = getFirstDateOfMonthlyCalendar(date);

    expect(firstDate).toEqual(new Date("2023/2/26"));
  });
});
