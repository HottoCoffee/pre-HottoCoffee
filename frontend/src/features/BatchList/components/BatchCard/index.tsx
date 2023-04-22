import { components } from "~/swagger/schema/schemas/batch";
import styles from "./BatchCard.module.scss";

type BatchType = components["schemas"]["Batch"];

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
