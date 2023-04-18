import { useRouter } from "next/router";
import { AiTwotoneHome } from "react-icons/ai";
import { MdViewList } from "react-icons/md";
import classNames from "classnames/bind";
import Link from "next/link";

import styles from "./SideNavigation.module.scss";

const clx = classNames.bind(styles);

export const SideNavigation = () => {
  const { pathname } = useRouter();

  const homeButtonClass = clx(styles.button, {
    [styles.active]: pathname === "/",
  });
  const batchListButtonClass = clx(styles.button, {
    [styles.active]: pathname === "/batch-list",
  });
  const containerClass = clx(styles.container);

  return (
    <div className={containerClass}>
      <Link href="/home" className={homeButtonClass}>
        <AiTwotoneHome />

        <span>Home</span>
      </Link>

      <Link href="/batch/list" className={batchListButtonClass}>
        <MdViewList />

        <span>Batch list</span>
      </Link>
    </div>
  );
};
