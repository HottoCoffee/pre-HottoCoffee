import * as Toast from "@radix-ui/react-toast";
import styles from "./Toaster.module.scss";

interface Props {
  type: "success" | "failed";
  description: React.ReactNode;
  title: string;
  open?: boolean;
  setOpen?: (open: boolean) => void;
}

export const Toaster = (props: Props) => {
  const { type, description, title, setOpen, open } = props;

  return (
    <>
      <Toast.Root className={styles.toastRoot} open={open} onOpenChange={setOpen}>
        <Toast.Title className={styles.toastTitle}>{title}</Toast.Title>
        <Toast.Description asChild>{description}</Toast.Description>
        <Toast.Action className={styles.toastAction} asChild altText="Remove message">
          <button className={styles.button}>Remove</button>
        </Toast.Action>
      </Toast.Root>
      <Toast.Viewport className={styles.toastViewport} />
    </>
  );
};
