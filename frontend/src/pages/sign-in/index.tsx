import { SignInUpForm } from "~/features/SignInUpForm";
import styles from "./index.module.scss";

export default function SignIn() {
  const onSuccess = (res: { token?: string }) => {
    console.log(res);
  };

  return (
    <div className={styles.container}>
      <SignInUpForm type="sign-in" onSuccess={onSuccess} />
    </div>
  );
}
