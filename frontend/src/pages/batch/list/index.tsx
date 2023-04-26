import { BatchList } from "~/features/BatchList";
import { DefaultLayout } from "~/features/Layouts/DefaultLayout";
import styles from "./index.module.scss";

export default function Home() {
  return (
    <DefaultLayout>
      <div className={styles.list}>
        <BatchList />
      </div>
    </DefaultLayout>
  );
}
