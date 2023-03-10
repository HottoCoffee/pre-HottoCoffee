import * as Form from "@radix-ui/react-form";
import { TextInput } from "~/shared/TextInput/ui";
import styles from "./RegisterNewBatchForm.modules.scss";
import classNames from "classnames/bind";
import { DatePickInput } from "~/shared/DatePicker/ui/DatePickInput";
import { FormEvent, useState } from "react";
import { createNewBatchValidation } from "../../utils/validations/createNewBatchValidation";
import { Button } from "~/shared/Button/ui";

const clx = classNames.bind(styles);

export const RegisterNewBatchForm = () => {
  const [selectedDate, setSelectedDate] = useState(new Date());

  const onSubmit = (event: FormEvent<HTMLFormElement>) => {
    event.preventDefault();
    const data = Object.fromEntries(new FormData(event.currentTarget));

    const validationResult = createNewBatchValidation({
      ...data,
      initial_date: selectedDate,
    });

    console.log(validationResult.data);
  };

  return (
    <Form.Root className={styles.fieldContainer} onSubmit={onSubmit}>
      <h1 className={styles.title}>Register New Batch</h1>
      <Form.Field name="batch_name" className={styles.formLayout}>
        <Form.FormLabel>Batch name</Form.FormLabel>
        <Form.Control asChild>
          <TextInput required type="text" />
        </Form.Control>

        <Form.Message match="valueMissing" className={styles.error}>
          Please enter batch name
        </Form.Message>
      </Form.Field>

      <Form.Field name="server_name" className={styles.formLayout}>
        <Form.FormLabel>Server name</Form.FormLabel>
        <Form.Control asChild>
          <TextInput required type="text" />
        </Form.Control>

        <Form.Message match="valueMissing" className={styles.error}>
          Please enter server name
        </Form.Message>
      </Form.Field>

      <Form.Field name="cron_setting" className={styles.formLayout}>
        <Form.FormLabel>Cron setting</Form.FormLabel>
        <Form.Control asChild>
          <TextInput required type="text" />
        </Form.Control>

        <Form.Message match="valueMissing" className={styles.error}>
          Please enter cron value
        </Form.Message>
      </Form.Field>

      <Form.Field name="initial_execution_date" className={styles.formLayout}>
        <Form.FormLabel>Initial execution date</Form.FormLabel>
        <Form.Control asChild>
          <DatePickInput
            selectedDate={selectedDate}
            onSelectDate={(date) => {
              setSelectedDate(date);
            }}
          />
        </Form.Control>
      </Form.Field>

      <Form.Field name="time_limit" className={styles.formLayout}>
        <Form.FormLabel>Time limit (s)</Form.FormLabel>
        <Form.Control asChild>
          <TextInput required type="number" />
        </Form.Control>

        <Form.Message match="valueMissing" className={styles.error}>
          Please enter valid number
        </Form.Message>
      </Form.Field>

      <div className={styles.footer}>
        <Button appearance="labeled" type="button">
          Back To Top
        </Button>

        <Form.Submit asChild>
          <Button appearance="success">Register</Button>
        </Form.Submit>
      </div>
    </Form.Root>
  );
};
