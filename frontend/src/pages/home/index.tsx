import { Header } from "~/shared/Header/ui";
import { CommonMetaInformation } from "~/shared/MetaInformation/CommonMetaInformation";
import styles from "./index.module.scss";
import { SideNavigation } from "~/shared/SideNavigation/ui";
import { MonthlyCalendar } from "~/features/MonthlyCalendar";
import { DateSwitcher } from "~/features/MonthlyCalendar/components/DateSwitcher";
import { useState } from "react";

export default function Home() {
  const [date, setDate] = useState(new Date());

  return (
    <>
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
    </>
  );
}
