export type Styles = {
  container: string;
  icon: string;
  input: string;
  popoverContent: string;
};

export type ClassNames = keyof Styles;

declare const styles: Styles;

export default styles;
