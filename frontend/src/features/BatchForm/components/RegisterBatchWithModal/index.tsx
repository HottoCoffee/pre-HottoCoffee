import styles from "./RegisterNewBatchWithModal.module.scss";
import * as Dialog from "@radix-ui/react-dialog";
import { components } from "~/swagger/schema/schemas/batch";
import { RegisterNewBatchForm } from "../RegisterBatch";
import { Button } from "~/shared/Button/ui";
import { useState } from "react";
import { Toaster } from "~/shared/Toaster/ui";
import * as Toast from "@radix-ui/react-toast";

interface Props {
  children: React.ReactNode;
}

export const RegisterNewBatchWithModal = (props: Props) => {
  const { children } = props;
  const [isOpen, setIsOpen] = useState(false);
  const [openToast, setOpenToast] = useState(false);

  const onSuccess = (batch: components["schemas"]["Batch"]) => {
    setIsOpen(false);
  };

  return (
    <div>
      <Dialog.Root open={isOpen} onOpenChange={setIsOpen}>
        <Dialog.Trigger asChild>{children}</Dialog.Trigger>
        <Dialog.Portal>
          <Dialog.Overlay className={styles.overlay} />
          <Dialog.Content className={styles.content}>
            <RegisterNewBatchForm
              onSuccess={onSuccess}
              footer={{
                closeButton: (
                  <Dialog.Close asChild>
                    <Button appearance="labeled" type="reset">
                      Close
                    </Button>
                  </Dialog.Close>
                ),
              }}
            />
          </Dialog.Content>
        </Dialog.Portal>
      </Dialog.Root>

      <Toaster type={"success"} description={<p>{""}</p>} title="Error on api." open={Boolean()} />
      <Toast.Viewport />
    </div>
  );
};
