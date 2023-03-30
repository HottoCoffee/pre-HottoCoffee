import classNames from "classnames/bind";
import { forwardRef, Ref } from "react";
import styles from "./Button.module.scss";

const clx = classNames.bind(styles);

interface Props extends Partial<JSX.IntrinsicElements["button"]> {
  className?: string;
  children: React.ReactNode;
  appearance: "default" | "danger" | "success" | "disabled" | "labeled";
}

export const Button = forwardRef(function Button(props: Props, ref: Ref<HTMLButtonElement>) {
  const { className, children, appearance, ...rest } = props;

  const classNameConcatenated = clx(
    styles.button,
    {
      [styles.default]: appearance === "default",
      [styles.danger]: appearance === "danger",
      [styles.success]: appearance === "success",
      [styles.disabled]: appearance === "disabled",
      [styles.labeled]: appearance === "labeled",
    },
    className,
  );

  return (
    <button className={classNameConcatenated} {...rest} ref={ref}>
      {children}
    </button>
  );
});
