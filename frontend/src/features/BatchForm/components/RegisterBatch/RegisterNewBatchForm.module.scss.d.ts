export type Styles = {
  error: string;
  fieldContainer: string;
  footer: string;
  formLayout: string;
  title: string;
};

export type ClassNames = keyof Styles;

declare const styles: Styles;

export default styles;
