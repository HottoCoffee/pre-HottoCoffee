import classNames from "classnames/bind";
import { forwardRef, Ref } from "react";
import styles from "./TextInput.modules.scss";

const clx = classNames.bind(styles);

interface Props extends Partial<JSX.IntrinsicElements["input"]> {
  className?: string;
}

export const TextInput = forwardRef(function TextInput(props: Props, ref: Ref<HTMLInputElement>) {
  const { className, ...rest } = props;

  const classNameConcatenated = clx(styles.input, className);

  return <input className={classNameConcatenated} {...rest} ref={ref} />;
});
