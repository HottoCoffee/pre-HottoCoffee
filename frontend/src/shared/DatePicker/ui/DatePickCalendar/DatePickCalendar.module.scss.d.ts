export type Styles = {
  arrowButton: string;
  container: string;
  day: string;
  displayedMonthLabel: string;
  week: string;
  wrapper: string;
};

export type ClassNames = keyof Styles;

declare const styles: Styles;

export default styles;
