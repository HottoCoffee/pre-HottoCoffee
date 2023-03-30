import { IoIosArrowForward, IoIosArrowBack } from "react-icons/io";
import * as AccessibleIcon from "@radix-ui/react-accessible-icon";
import { dateSwitcherOnChangeMethod } from "./functions";
import { format } from "date-fns";
import styles from "./DateSwitcher.modules.scss";

/**
 * @package
 */
export interface Props {
  date: Date;
  onChange: (date: Date) => void;
  type: "month" | "week";
}

export const DateSwitcher = (props: Props) => {
  const { date, onChange } = props;

  const { goNext, backPrevious } = dateSwitcherOnChangeMethod(props);

  return (
    <div className={styles.container}>
      <button className={styles.button} onClick={() => onChange(backPrevious())}>
        <AccessibleIcon.Root label="bak to previous month">
          <IoIosArrowBack />
        </AccessibleIcon.Root>
      </button>

      <p>{format(date, "yyyy/MM")}</p>

      <button className={styles.button} onClick={() => onChange(goNext())}>
        <AccessibleIcon.Root label="Go to next month">
          <IoIosArrowForward />
        </AccessibleIcon.Root>
      </button>
    </div>
  );
};
