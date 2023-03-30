import { add, endOfMonth, getWeek, startOfMonth, startOfWeek } from "date-fns";

export const getDateListInMonth = (date: Date) => {
  const firstDay = getFirstDateOfMonthlyCalendar(date);
  const weekCount = getWeekCount(date);

  return [...new Array(weekCount)].map((_, weekIndex) => {
    return [...new Array(7)].map((_, dayIndex) => {
      return add(firstDay, { weeks: weekIndex, days: dayIndex });
    });
  });
};

/**
 *
 * @package
 * @param date
 * @returns the first date of monthly calendar
 *
 * @example (date: 2023/3) -> return: 2023/2/29
 *
 */
export const getFirstDateOfMonthlyCalendar = (date: Date) => {
  return startOfWeek(startOfMonth(date));
};

/**
 * return week count in provided month
 * @param date
 * @returns number
 *
 * @example
 * const date = new Date("2023/3/1");
 * const count = getWeekCount(date); // -> 5
 */
export const getWeekCount = (date: Date): number => {
  const firstDate = startOfMonth(date);
  const lastDate = endOfMonth(date);

  let currentMonthWeeks = getWeek(lastDate) - getWeek(firstDate) + 1;

  if (currentMonthWeeks < 0) {
    currentMonthWeeks += getLastWeekNumberInLastYear(firstDate);
  }

  return currentMonthWeeks;
};

/**
 * Get last week number in last year
 */
const getLastWeekNumberInLastYear = (date: Date): number => {
  const currentYear = date.getFullYear();
  const lastDayInLastYear = new Date(currentYear - 1, 11, 31);
  let weekNumber = getWeek(lastDayInLastYear);

  if (weekNumber === 1) {
    weekNumber = getWeek(add(lastDayInLastYear, { weeks: -1 }));
  }

  return weekNumber;
};
