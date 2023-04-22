export type Styles = {
  batchName: string;
  frame: string;
  serverName: string;
};

export type ClassNames = keyof Styles;

declare const styles: Styles;

export default styles;
