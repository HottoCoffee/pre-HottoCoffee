import { SignInUpForm } from "~/features/SignInUpForm";
import styles from "./index.module.scss";
import { useRouter } from "next/router";

export default function SignUp() {
  const router = useRouter();

  const onSuccess = () => {
    router.push("/sign-in");
  };

  return (
    <div className={styles.container}>
      <SignInUpForm type="sign-up" onSuccess={onSuccess} />
    </div>
  );
}
