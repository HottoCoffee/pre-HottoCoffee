import { components } from "~/swagger/schema/schemas/history";
import { BatchStatusButton } from "../BatchStatusButton";
import styles from "./MonthlyCalendarDayCell.module.scss";
import { format } from "date-fns";
import { useElementSize } from "usehooks-ts";
import { calculateRenderTarget } from "./functions";
import * as Popover from "@radix-ui/react-popover";
import { AiOutlineCloseCircle } from "react-icons/ai";
import classNames from "classnames/bind";

type History = components["schemas"]["History"];

const clx = classNames.bind(styles);

interface Props {
  date: Date;
  batchHistoryList: History[];
}

export const MonthlyCalendarDayCell = (props: Props) => {
  const { batchHistoryList, date } = props;
  const [ref, { height }] = useElementSize();

  const { hasMore, showCount, restCount } = calculateRenderTarget(batchHistoryList.length, height);
  const showBatchHistoryList = hasMore ? batchHistoryList.slice(0, showCount) : batchHistoryList;
  const restBatchHistoryList = hasMore ? batchHistoryList.slice(showCount) : [];

  const restBatchList = clx(styles.list, styles.restBatchList);

  return (
    <div ref={ref} className={styles.container}>
      <p>{format(date, "d")}</p>

      <ul className={styles.list}>
        {showBatchHistoryList.map((batch, index) => {
          return (
            <li key={batch.batch_id + ":" + index}>
              <BatchStatusButton
                batchName={batch.batch_name}
                batchId={batch.batch_id}
                status={batch.status}
                historyId={batch.history_id}
              />
            </li>
          );
        })}
      </ul>

      {hasMore && (
        <Popover.Root>
          <Popover.Trigger asChild>
            <button className={styles.moreButton}>Show more {restCount}</button>
          </Popover.Trigger>

          <Popover.Portal>
            <Popover.Content className={styles.popover}>
              <Popover.Close className={styles.close}>
                <AiOutlineCloseCircle />
              </Popover.Close>

              <ul className={restBatchList}>
                {restBatchHistoryList.map((batch, index) => {
                  return (
                    <li key={batch.batch_id + ":" + index}>
                      <BatchStatusButton
                        batchName={batch.batch_name}
                        batchId={batch.batch_id}
                        status={batch.status}
                        historyId={batch.history_id}
                      />
                    </li>
                  );
                })}
              </ul>
            </Popover.Content>
          </Popover.Portal>
        </Popover.Root>
      )}
    </div>
  );
};
