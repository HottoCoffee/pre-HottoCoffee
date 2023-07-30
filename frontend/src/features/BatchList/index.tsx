import { useAspidaQuery } from "@aspida/react-query";
import { client } from "~/modules/aspidaClient";
import { BatchCard } from "./components/BatchCard";
import styles from "./BatchList.module.scss";
import { NewBatchCard } from "./components/NewBatchCard";
import Skeleton from "react-loading-skeleton";
import { TextInput } from "~/shared/TextInput/ui";
import { useState } from "react";
import { useDebounce } from "usehooks-ts";
import { useUserInformation } from "~/hooks/useUserInformation";

const skeletons: number[] = [...new Array(10)].fill(0);

export const BatchList = () => {
  const [keyword, setKeyword] = useState("");
  const debounced = useDebounce(keyword, 500);
  const { workspaceId } = useUserInformation();

  const { data, isLoading } = useAspidaQuery(
    client.api.workspace._workspace_id(workspaceId).batch,
    {
      query: {
        keyword: debounced,
      },
      keepPreviousData: true,
    },
  );

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
    <div>
      <div className={styles.inputWrapper}>
        <TextInput onChange={(e) => setKeyword(e.target.value)} value={keyword} />
      </div>

      <ul className={styles.list}>
        {data?.map((batch) => {
          return (
            <li key={batch.id}>
              <BatchCard batch={batch} />
            </li>
          );
        })}

        <li>
          <NewBatchCard />
        </li>
      </ul>
    </div>
  );
};
