export type Styles = {
  close: string;
  container: string;
  list: string;
  moreButton: string;
  popover: string;
  restBatchList: string;
  slideDownAndFade: string;
  slideLeftAndFade: string;
  slideRightAndFade: string;
  slideUpAndFade: string;
};

export type ClassNames = keyof Styles;

declare const styles: Styles;

export default styles;
