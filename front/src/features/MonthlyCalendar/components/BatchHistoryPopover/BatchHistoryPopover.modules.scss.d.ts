export type Styles = {
  container: string;
  errorText: string;
  failed: string;
  informationContainer: string;
  inProgress: string;
  separator: string;
  success: string;
};

export type ClassNames = keyof Styles;

declare const styles: Styles;

export default styles;
