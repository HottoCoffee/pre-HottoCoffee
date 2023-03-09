export type Styles = {
  container: string;
  errorText: string;
  informationContainer: string;
  separator: string;
};

export type ClassNames = keyof Styles;

declare const styles: Styles;

export default styles;
