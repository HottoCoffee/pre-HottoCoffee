export type Styles = {
  contentShow: string;
  overlayShow: string;
  slideDownAndFade: string;
  slideLeftAndFade: string;
  slideRightAndFade: string;
  slideUpAndFade: string;
};

export type ClassNames = keyof Styles;

declare const styles: Styles;

export default styles;
