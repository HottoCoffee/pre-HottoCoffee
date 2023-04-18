export type Styles = {
  container: string;
  main: string;
  registerNewBatch: string;
};

export type ClassNames = keyof Styles;

declare const styles: Styles;

export default styles;
