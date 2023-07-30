import { DefaultLayout } from "~/features/Layouts/DefaultLayout";
import styles from "./index.module.scss";
import { GetServerSideProps } from "next";
import { useAspidaQuery } from "@aspida/react-query";
import { client } from "~/modules/aspidaClient";
import { EditBatchForm } from "~/features/BatchForm/components/EditBatchForm";
import { useRouter } from "next/router";
import { useUserInformation } from "~/hooks/useUserInformation";

interface Props {
  id: number;
}

export default function Home(props: Props) {
  const { workspaceId } = useUserInformation();
  const { data, isLoading } = useAspidaQuery(
    client.api.workspace._workspace_id(workspaceId).batch._batch_id(props.id),
    {
      enabled: Boolean(workspaceId),
    },
  );
  const router = useRouter();

  const onSuccess = () => {
    router.push("/batch/list");
  };

  if (isLoading || !data) {
    return <div className={styles.container}>loading...</div>;
  }

  return (
    <DefaultLayout>
      <div className={styles.container}>
        <EditBatchForm initialBatch={data} onSuccess={onSuccess} />
      </div>
    </DefaultLayout>
  );
}

export const getServerSideProps: GetServerSideProps = async (context) => {
  const { id } = context.query;
  const parsedAsNumber = Number(id);

  if (!parsedAsNumber) {
    return {
      notFound: true,
    };
  }

  return {
    props: { id: parsedAsNumber },
  };
};
