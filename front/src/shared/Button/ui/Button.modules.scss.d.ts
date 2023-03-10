export type Styles = {
  button: string;
  danger: string;
  default: string;
  disabled: string;
  labeled: string;
  success: string;
};

export type ClassNames = keyof Styles;

declare const styles: Styles;

export default styles;
