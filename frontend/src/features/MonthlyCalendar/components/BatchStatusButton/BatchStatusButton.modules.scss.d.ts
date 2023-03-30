export type Styles = {
  beforeStarted: string;
  close: string;
  container: string;
  failed: string;
  inProgress: string;
  loadingIcon: string;
  popover: string;
  slideDownAndFade: string;
  slideLeftAndFade: string;
  slideRightAndFade: string;
  slideUpAndFade: string;
  success: string;
};

export type ClassNames = keyof Styles;

declare const styles: Styles;

export default styles;
