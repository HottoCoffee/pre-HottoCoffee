import { useAspidaQuery } from "@aspida/react-query";
import Skeleton from "react-loading-skeleton";
import { client } from "~/modules/aspidaClient";
import classNames from "classnames/bind";

const clx = classNames.bind(styles);

import styles from "./BatchHistoryPopover.module.scss";
import { format } from "date-fns";

interface Props {
  batchId: number;
  historyId: number;
}

export const BatchHistoryPopover = (props: Props) => {
  const { batchId, historyId } = props;

  const { data, error } = useAspidaQuery(
    client.api.batch._batch_id(batchId).history._history_id(historyId),
  );

  const informationContainerClass = clx(styles.container, styles.informationContainer);
  const statusTextClass = clx({
    [styles.failed]: data?.status === "failed",
    [styles.success]: data?.status === "success",
    [styles.inProgress]: data?.status === "in_progress",
  });

  if (error) {
    return <p className={styles.errorText}>Unknown Error occurred.</p>;
  }

  if (!data) {
    return (
      <div className={styles.container}>
        <Skeleton count={3} />
      </div>
    );
  }

  return (
    <div className={informationContainerClass}>
      <h3>{data.batch_name}</h3>
      <span />

      <p className={styles.separator}>Status</p>
      <p className={statusTextClass}>{data.status.split("_").join(" ")}</p>

      <p className={styles.separator}>start</p>
      <p className="">{format(new Date(data.start_datetime), "yyyy/MM/dd HH:mm:ss")}</p>
    </div>
  );
};
