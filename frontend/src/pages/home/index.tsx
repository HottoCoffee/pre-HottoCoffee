import { Header } from "~/shared/Header/ui";
import { CommonMetaInformation } from "~/shared/MetaInformation/CommonMetaInformation";
import styles from "./index.module.scss";
import { SideNavigation } from "~/shared/SideNavigation/ui";
import { MonthlyCalendar } from "~/features/MonthlyCalendar";
import { DateSwitcher } from "~/features/MonthlyCalendar/components/DateSwitcher";
import { useState } from "react";
import { RegisterNewBatchWithModal } from "~/features/BatchForm/components/RegisterBatchWithModal";
import * as AccessibleIcon from "@radix-ui/react-accessible-icon";

import { AiOutlinePlus } from "react-icons/ai";

export default function Home() {
  const [date, setDate] = useState(new Date());

  return (
    <>
      <div className={styles.container}>
        <CommonMetaInformation />

        <Header>
          <DateSwitcher date={date} onChange={setDate} type={"month"} />
        </Header>

        <main className={styles.main}>
          <SideNavigation />

          <div>
            <MonthlyCalendar date={date} />
          </div>
        </main>
      </div>

      <RegisterNewBatchWithModal>
        <button className={styles.registerNewBatch}>
          <AccessibleIcon.Root label="register new batch">
            <AiOutlinePlus size={32} fill="white" />
          </AccessibleIcon.Root>
        </button>
      </RegisterNewBatchWithModal>
    </>
  );
}
