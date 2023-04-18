export type Styles = {
  content: string;
  contentShow: string;
  overlay: string;
  overlayShow: string;
  slideDownAndFade: string;
  slideLeftAndFade: string;
  slideRightAndFade: string;
  slideUpAndFade: string;
};

export type ClassNames = keyof Styles;

declare const styles: Styles;

export default styles;
