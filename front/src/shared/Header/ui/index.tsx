import styles from "./Header.modules.scss";

interface Props {
  children: React.ReactNode;
}

/**
 * @public
 */
export const Header = (props: Props) => {
  const { children } = props;

  return (
    <div className={styles.container}>
      <h1 className={styles.title}>HottoCoffee</h1>

      <div>{children}</div>
    </div>
  );
};
