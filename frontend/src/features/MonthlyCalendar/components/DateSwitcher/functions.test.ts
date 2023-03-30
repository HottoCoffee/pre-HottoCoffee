import { dateSwitcherOnChangeMethod } from "./functions";

describe("features > MonthlyCalenadr > components > DateSwitcher > functions", () => {
  test("returns incremented 1 month when user calls goNext method with type='month'", () => {
    const { goNext } = dateSwitcherOnChangeMethod({
      date: new Date("2023/3/1"),
      type: "month",
      onChange: jest.fn(),
    });

    expect(goNext()).toEqual(new Date("2023/4/1"));
  });
  test("returns decremented 1 month when user calls backPrevious method with type='month'", () => {
    const { backPrevious } = dateSwitcherOnChangeMethod({
      date: new Date("2023/3/1"),
      type: "month",
      onChange: jest.fn(),
    });

    expect(backPrevious()).toEqual(new Date("2023/2/1"));
  });

  test("returns decremented 1 week when user calls goNext method with type='week'", () => {
    const { goNext } = dateSwitcherOnChangeMethod({
      date: new Date("2023/3/1"),
      type: "week",
      onChange: jest.fn(),
    });

    expect(goNext()).toEqual(new Date("2023/3/8"));
  });
  test("returns decremented 1 week when user calls backPrevious method with type='week'", () => {
    const { backPrevious } = dateSwitcherOnChangeMethod({
      date: new Date("2023/3/1"),
      type: "week",
      onChange: jest.fn(),
    });

    expect(backPrevious()).toEqual(new Date("2023/2/22"));
  });
});
