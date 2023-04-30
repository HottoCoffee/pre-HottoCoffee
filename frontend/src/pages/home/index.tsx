import styles from "./index.module.scss";
import { MonthlyCalendar } from "~/features/MonthlyCalendar";
import { DateSwitcher } from "~/features/MonthlyCalendar/components/DateSwitcher";
import { useState } from "react";
import { RegisterNewBatchWithModal } from "~/features/BatchForm/components/RegisterBatchWithModal";
import * as AccessibleIcon from "@radix-ui/react-accessible-icon";

import { AiOutlinePlus } from "react-icons/ai";
import { DefaultLayout } from "~/features/Layouts/DefaultLayout";

export default function Home() {
  const [date, setDate] = useState(new Date());

  return (
    <DefaultLayout headerChildren={<DateSwitcher date={date} onChange={setDate} type={"month"} />}>
      <MonthlyCalendar date={date} />

      <RegisterNewBatchWithModal>
        <button className={styles.registerNewBatch}>
          <AccessibleIcon.Root label="register new batch">
            <AiOutlinePlus size={32} fill="white" />
          </AccessibleIcon.Root>
        </button>
      </RegisterNewBatchWithModal>
    </DefaultLayout>
  );
}
