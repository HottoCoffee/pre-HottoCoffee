import * as Form from "@radix-ui/react-form";
import { TextInput } from "~/shared/TextInput/ui";

import styles from "./index.module.scss";
import Link from "next/link";
import { Button } from "~/shared/Button/ui";
import { FormEventHandler, useState } from "react";
import { Toaster } from "~/shared/Toaster/ui";
import { client } from "~/modules/aspidaClient";
import * as Toast from "@radix-ui/react-toast";
import { signInUpFormValidation } from "./validation";
import { isAxiosError } from "axios";
import { components } from "~/swagger/schema/schemas/success";

interface SignInProps {
  type: "sign-in";
  onSuccess: (res: { token?: string }) => void;
}
interface SignUpProps {
  type: "sign-up";
  onSuccess: (res: components["schemas"]["Success"]) => void;
}

type Props = SignInProps | SignUpProps;

export const SignInUpForm = (props: Props) => {
  const { type, onSuccess } = props;
  const [errorMessage, setErrorMessage] = useState<string | undefined>();

  const onSubmit: FormEventHandler<HTMLFormElement> = async (e) => {
    e.preventDefault();
    const data = Object.fromEntries(new FormData(e.currentTarget));

    const validationResult = signInUpFormValidation({
      ...data,
    });

    if (validationResult.type === "err") {
      window.alert("Please input valid values");
      return;
    }

    try {
      if (type === "sign-up") {
        const res = await client.api.public.sign_up.post({
          body: validationResult.data,
        });
        return onSuccess(res.body);
      }
      if (type === "sign-in") {
        const res = await client.api.public.sign_in.post({
          body: validationResult.data,
        });
        return onSuccess(res.body);
      }
    } catch (e) {
      if (isAxiosError(e)) {
        setErrorMessage(e.message);
      }
    }
  };

  const clearError = () => {
    setErrorMessage(undefined);
  };

  return (
    <>
      <div className={styles.container}>
        <header className={styles.header}>
          <h1 className={styles.title}>
            HottoCoffee - {type === "sign-in" ? "Sign In" : "Sign Up"}
          </h1>

          <Link href={type === "sign-in" ? "/sign-up" : "/sign-in"} className={styles.link}>
            or {type === "sign-in" ? "Sign Up" : "Sign In"}
          </Link>
        </header>

        <Form.Root className={styles.form} onSubmit={onSubmit}>
          <Form.Field name="email" className={styles.field}>
            <Form.FormLabel>email</Form.FormLabel>
            <Form.Control asChild>
              <TextInput required type="email" placeholder="example@example.com" />
            </Form.Control>

            <Form.Message match="valueMissing" className={styles.error}>
              Please enter email
            </Form.Message>

            <Form.Message match="typeMismatch" className={styles.error}>
              Please enter valid email
            </Form.Message>
          </Form.Field>

          <Form.Field name="password" className={styles.field}>
            <Form.FormLabel>password</Form.FormLabel>
            <Form.Control asChild>
              <TextInput required type="password" />
            </Form.Control>

            <Form.Message match="valueMissing" className={styles.error}>
              Please enter password
            </Form.Message>
          </Form.Field>

          <Form.Submit asChild>
            <Button type="submit" appearance="success">
              {type === "sign-in" ? "Sign In" : "Sign Up"}
            </Button>
          </Form.Submit>
        </Form.Root>
      </div>

      <Toaster
        type="failed"
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
    </>
  );
};
