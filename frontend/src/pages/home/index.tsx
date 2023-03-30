import { Header } from "~/shared/Header/ui";
import { CommonMetaInformation } from "~/shared/MetaInformation/CommonMetaInformation";
import styles from "./index.module.scss";
import { SideNavigation } from "~/shared/SideNavigation/ui";

export default function Home() {
  return (
    <>
      <CommonMetaInformation />

      <Header />

      <body className={styles.body}>
        <SideNavigation />

        <div></div>
      </body>
    </>
  );
}
