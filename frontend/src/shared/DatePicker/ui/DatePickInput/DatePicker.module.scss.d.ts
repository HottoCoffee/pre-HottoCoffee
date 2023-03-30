export type Styles = {
  button: string;
  container: string;
  icon: string;
  placeholder: string;
  popoverContent: string;
};

export type ClassNames = keyof Styles;

declare const styles: Styles;

export default styles;
