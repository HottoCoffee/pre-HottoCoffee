export type Styles = {
  container: string;
  error: string;
  field: string;
  form: string;
  header: string;
  link: string;
  title: string;
};

export type ClassNames = keyof Styles;

declare const styles: Styles;

export default styles;
