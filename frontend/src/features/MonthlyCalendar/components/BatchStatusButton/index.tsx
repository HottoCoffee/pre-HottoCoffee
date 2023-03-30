import { HottoCoffee } from "~/types";
import classNames from "classnames/bind";
import * as AccessibleIcon from "@radix-ui/react-accessible-icon";
import { ThreeDots } from "react-loading-icons";
import * as Popover from "@radix-ui/react-popover";
import { AiOutlineCloseCircle } from "react-icons/ai";
import styles from "./BatchStatusButton.module.scss";
import { BatchDetailPopover } from "../BatchDetailPopover";
import { BatchHistoryPopover } from "../BatchHistoryPopover";
import { HEIGHT_OF_BATCH_STATUS_BUTTON } from "./constants";

const clx = classNames.bind(styles);

interface Props {
  status: HottoCoffee.BatchStatus;
  batchName: string;
  batchId: number;
  historyId?: number;
}

export const BatchStatusButton = (props: Props) => {
  const { batchName, status, batchId, historyId } = props;

  const buttonClass = clx(styles.container, {
    [styles.beforeStarted]: status === "before_started",
    [styles.inProgress]: status === "in_progress",
    [styles.success]: status === "success",
    [styles.failed]: status === "failed",
  });

  return (
    <Popover.Root>
      <Popover.Trigger asChild>
        <button className={buttonClass} style={{ height: HEIGHT_OF_BATCH_STATUS_BUTTON }}>
          {batchName}
          {status === "in_progress" && (
            <AccessibleIcon.Root label="Loading effect">
              <ThreeDots className={styles.loadingIcon} fill="#373737" />
            </AccessibleIcon.Root>
          )}
        </button>
      </Popover.Trigger>

      <Popover.Portal>
        <Popover.Content className={styles.popover}>
          <div>
            <Popover.Close className={styles.close}>
              <AiOutlineCloseCircle />
            </Popover.Close>

            {status === "before_started" && <BatchDetailPopover batchId={batchId} />}
            {status !== "before_started" && historyId && (
              <BatchHistoryPopover batchId={batchId} historyId={historyId} />
            )}
          </div>
        </Popover.Content>
      </Popover.Portal>
    </Popover.Root>
  );
};
