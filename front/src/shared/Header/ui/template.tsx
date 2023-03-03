import styles from "./style.modules.scss";

interface Props {
  children: React.ReactNode;
}

export const Template = (props: Props) => {
  const { children } = props;

  return (
    <div className={styles.container}>
      <h1 className={styles.title}>HottoCoffee</h1>

      <div>{children}</div>
    </div>
  );
};
