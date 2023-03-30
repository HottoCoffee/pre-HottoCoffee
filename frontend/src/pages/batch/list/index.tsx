import Link from "next/link";
import { Header } from "~/shared/Header/ui";
import { CommonMetaInformation } from "~/shared/MetaInformation/CommonMetaInformation";

export default function Home() {
  return (
    <>
      <CommonMetaInformation />

      <Header />

      <main>
        <p>Under construction......</p>
        <br />
        <br />
        <br />
        <br />
        <br />

        <Link href="/home">Back to top</Link>
      </main>
    </>
  );
}
