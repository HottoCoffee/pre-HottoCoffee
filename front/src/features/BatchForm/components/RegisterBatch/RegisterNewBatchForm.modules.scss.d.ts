export type Styles = {
  error: string;
  fieldContainer: string;
  formLayout: string;
};

export type ClassNames = keyof Styles;

declare const styles: Styles;

export default styles;
