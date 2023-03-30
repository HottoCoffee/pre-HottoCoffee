export type Styles = {
  button: string;
  hide: string;
  slideIn: string;
  swipeOut: string;
  toastAction: string;
  ToastDescription: string;
  toastRoot: string;
  toastTitle: string;
  toastViewport: string;
};

export type ClassNames = keyof Styles;

declare const styles: Styles;

export default styles;
