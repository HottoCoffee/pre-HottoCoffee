import { AiOutlinePlus } from "react-icons/ai";
import * as AccessibleIcon from "@radix-ui/react-accessible-icon";
import styles from "./NewBatchCard.module.scss";
import { RegisterNewBatchWithModal } from "~/features/BatchForm/components/RegisterBatchWithModal";

export const NewBatchCard = () => {
  return (
    <RegisterNewBatchWithModal>
      <button className={styles.button}>
        <AccessibleIcon.Root label="create new batch">
          <AiOutlinePlus size={28} />
        </AccessibleIcon.Root>

        <span className={styles.text}>New</span>
      </button>
    </RegisterNewBatchWithModal>
  );
};
