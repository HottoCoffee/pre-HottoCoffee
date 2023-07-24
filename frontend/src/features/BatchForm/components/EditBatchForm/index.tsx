import * as Form from "@radix-ui/react-form";
import { TextInput } from "~/shared/TextInput/ui";
import styles from "./EditBatchForm.module.scss";
import { DatePickInput } from "~/shared/DatePicker/ui/DatePickInput";
import { FormEvent, useState } from "react";
import { Button } from "~/shared/Button/ui";
import { client } from "~/modules/aspidaClient";
import * as Toast from "@radix-ui/react-toast";
import { components } from "~/swagger/schema/schemas/batch";
import axios from "axios";
import { Toaster } from "~/shared/Toaster/ui";
import { editBatchValidation } from "../../utils/validations/editNewBatchValidation";
import { useUserInformation } from "~/hooks/useUserInformation";

interface Props {
  onSuccess: (batch: components["schemas"]["Batch"]) => void;
  footer?: {
    closeButton?: React.ReactNode;
  };
  initialBatch: components["schemas"]["Batch"];
}

export const EditBatchForm = (props: Props) => {
  const { onSuccess, footer, initialBatch } = props;
  const [selectedDate, setSelectedDate] = useState(new Date(initialBatch.initial_date));
  const [errorMessage, setErrorMessage] = useState<string | undefined>();
  const { workspaceId } = useUserInformation();

  const onSubmit = async (event: FormEvent<HTMLFormElement>) => {
    event.preventDefault();
    const data = Object.fromEntries(new FormData(event.currentTarget));

    const validationResult = editBatchValidation({
      ...data,
      initial_date: selectedDate,
      id: initialBatch.id,
    });

    if (validationResult.type === "err") {
      window.alert("Please input valid values.");
      return;
    }

    try {
      const response = await client.api.workspace
        ._workspace_id(workspaceId)
        .batch._batch_id(initialBatch.id)
        .put({
          body: validationResult.data,
        });

      onSuccess?.(response.body);
    } catch (e) {
      if (axios.isAxiosError(e)) {
        setErrorMessage(e.message);
      }
    }
  };

  const clearError = () => {
    setErrorMessage(undefined);
  };

  return (
    <Form.Root className={styles.fieldContainer} onSubmit={onSubmit}>
      <h1 className={styles.title}>Edit {initialBatch.batch_name}</h1>
      <Form.Field name="batch_name" className={styles.formLayout}>
        <Form.FormLabel>Batch name</Form.FormLabel>
        <Form.Control asChild>
          <TextInput required type="text" defaultValue={initialBatch.batch_name} />
        </Form.Control>

        <Form.Message match="valueMissing" className={styles.error}>
          Please enter batch name
        </Form.Message>
      </Form.Field>

      <Form.Field name="server_name" className={styles.formLayout}>
        <Form.FormLabel>Server name</Form.FormLabel>
        <Form.Control asChild>
          <TextInput required type="text" defaultValue={initialBatch.server_name} />
        </Form.Control>

        <Form.Message match="valueMissing" className={styles.error}>
          Please enter server name
        </Form.Message>
      </Form.Field>

      <Form.Field name="cron_setting" className={styles.formLayout}>
        <Form.FormLabel>Cron setting</Form.FormLabel>
        <Form.Control asChild>
          <TextInput required type="text" defaultValue={initialBatch.cron_setting} />
        </Form.Control>

        <Form.Message match="valueMissing" className={styles.error}>
          Please enter cron value
        </Form.Message>
      </Form.Field>

      <Form.Field name="initial_execution_date" className={styles.formLayout}>
        <Form.FormLabel>Initial execution date</Form.FormLabel>
        <Form.Control id="initial_execution_date" asChild>
          <DatePickInput
            selectedDate={selectedDate}
            onSelectDate={(date) => {
              setSelectedDate(date);
            }}
            dataTestId="initial_execution_date"
          />
        </Form.Control>
      </Form.Field>

      <Form.Field name="time_limit" className={styles.formLayout}>
        <Form.FormLabel>Time limit (min)</Form.FormLabel>
        <Form.Control asChild>
          <TextInput required type="number" defaultValue={initialBatch.time_limit} />
        </Form.Control>

        <Form.Message match="valueMissing" className={styles.error}>
          Please enter valid number
        </Form.Message>
      </Form.Field>

      <div className={styles.footer}>
        {footer?.closeButton ?? (
          <Button appearance="labeled" type="reset">
            Back
          </Button>
        )}

        <Form.Submit asChild>
          <Button type="submit" appearance="success" data-testid="submit">
            Update
          </Button>
        </Form.Submit>
      </div>

      <Toaster
        type={"success"}
        description={<p>{errorMessage}</p>}
        title="Error on api."
        open={Boolean(errorMessage)}
        setOpen={(open: boolean) => {
          if (!open) {
            clearError();
          }
        }}
      />
      <Toast.Viewport />
    </Form.Root>
  );
};
