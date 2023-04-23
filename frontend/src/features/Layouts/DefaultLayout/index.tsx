import { CommonMetaInformation } from "~/shared/MetaInformation/CommonMetaInformation";
import { SideNavigation } from "~/shared/SideNavigation/ui";
import styles from "./DefaultLayout.module.scss";
import { Header } from "~/shared/Header/ui";

interface Props {
  children: React.ReactNode;
  headerChildren?: React.ReactNode;
}

export const DefaultLayout = (props: Props) => {
  const { children, headerChildren } = props;

  return (
    <div className={styles.container}>
      <CommonMetaInformation />

      <Header>{headerChildren}</Header>

      <main className={styles.main}>
        <SideNavigation />

        <div>{children}</div>
      </main>
    </div>
  );
};
