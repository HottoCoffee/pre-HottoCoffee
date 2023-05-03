import { components } from "~/swagger/schema/schemas/batch";
import styles from "./BatchCard.module.scss";
import { useRouter } from "next/router";

type BatchType = components["schemas"]["Batch"];

interface Props {
  batch: BatchType;
}

export const BatchCard = (props: Props) => {
  const { batch_name, server_name, id } = props.batch;
  const router = useRouter();

  const onClick = () => {
    router.push(`/batch/edit/${id}`);
  };

  return (
    <button className={styles.frame} onClick={onClick}>
      <p className={styles.batchName}>{batch_name}</p>
      <p className={styles.serverName}>server: {server_name}</p>
    </button>
  );
};
