import { useAspidaQuery } from "@aspida/react-query";
import { client } from "~/modules/aspidaClient";
import { BatchCard } from "./components/BatchCard";
import styles from "./BatchList.module.scss";
import { NewBatchCard } from "./components/NewBatchCard";
import Skeleton from "react-loading-skeleton";

const skeletons: number[] = [...new Array(10)].fill(0);

export const BatchList = () => {
  const { data, isLoading } = useAspidaQuery(client.api.batch);

  if (isLoading && !data) {
    return (
      <ul className={styles.list}>
        {skeletons.map((_, index) => {
          return <Skeleton width={220} height={80} key={`skelton-${index}`} />;
        })}
      </ul>
    );
  }

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