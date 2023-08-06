import { z } from "zod";
import { Result } from "~/types/Result";

const schema = z.object({
  email: z.string().email(),
  password: z.string(),
});

export const signInUpFormValidation = (formData: {
  [key: string]: FormDataEntryValue;
}): Result<
  {
    email: string;
    password: string;
  },
  z.ZodError<object>
> => {
  const result = schema.safeParse(formData);
  if (result.success) {
    return { type: "ok", data: result.data };
  } else {
    return { type: "err", data: result.error };
  }
};
