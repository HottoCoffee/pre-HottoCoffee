import { useAspidaQuery } from "@aspida/react-query";
import Skeleton from "react-loading-skeleton";
import { client } from "~/modules/aspidaClient";
import classNames from "classnames/bind";

const clx = classNames.bind(styles);

import styles from "./BatchDetailPopover.modules.scss";

interface Props {
  batchId: number;
}

export const BatchDetailPopover = (props: Props) => {
  const { batchId } = props;

  const { data, error } = useAspidaQuery(client.api.batch._batch_id(batchId));

  const informationContainerClass = clx(styles.container, styles.informationContainer);

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

      <p className={styles.separator}>Server Name</p>
      <p>{data.server_name}</p>

      <p className={styles.separator}>cron setting</p>
      <p>{data.cron_setting}</p>

      <p className={styles.separator}>time limit</p>
      <p>{data.time_limit} (min)</p>
    </div>
  );
};
