import { components } from "~/swagger/schema/schemas/batch";
import styles from "./BatchCard.module.scss";

type BatchType = components["schemas"]["Batch"];
const batchMock: BatchType = {
  id: 1,
  batch_name: "エキスポート",
  server_name: "初号機",
  cron_setting: "0 0 * * *",
  initial_date: "2023-01-01",
  time_limit: 60,
};

interface Props {
  batch: BatchType;
}

export const BatchCard = (props: Props) => {
  const { batch_name, server_name } = props.batch;

  return (
    <button className={styles.frame}>
      <p className={styles.batchName}>{batch_name}</p>
      <p className={styles.serverName}>server: {server_name}</p>
    </button>
  );
};
