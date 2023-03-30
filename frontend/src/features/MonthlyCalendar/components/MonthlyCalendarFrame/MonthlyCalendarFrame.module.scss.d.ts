export type Styles = {
  container: string;
  day: string;
  week: string;
};

export type ClassNames = keyof Styles;

declare const styles: Styles;

export default styles;
