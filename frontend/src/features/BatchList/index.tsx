import { useAspidaQuery } from "@aspida/react-query";
import { client } from "~/modules/aspidaClient";
import { BatchCard } from "./components/BatchCard";
import styles from "./BatchList.module.scss";
import { NewBatchCard } from "./components/NewBatchCard";

export const BatchList = () => {
  const { data } = useAspidaQuery(client.api.batch);

  return (
    <ul className={styles.list}>
      {data?.map((batch) => {
        return (
          <li key={`batch-card-${batch.id}`}>
            <BatchCard batch={batch} />
          </li>
        );
      })}

      <li>
        <NewBatchCard />
      </li>
    </ul>
  );
};
