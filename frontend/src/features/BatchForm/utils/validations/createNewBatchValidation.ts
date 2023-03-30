import { format } from "date-fns";
import z from "zod";
import type { components } from "~/swagger/schema/schemas/batch";
import { Result } from "~/types/Result";

type CreateNewBatchValidationSchema = Omit<components["schemas"]["Batch"], "id">;
const createNewBatchValidationSchema: z.ZodSchema<
  CreateNewBatchValidationSchema,
  z.ZodTypeDef,
  object
> = z.object({
  batch_name: z.string().min(1),
  server_name: z.string().min(1),
  cron_setting: z.string().min(1),
  initial_date: z.date().transform((date) => {
    return date.toISOString();
  }),
  time_limit: z.preprocess(Number, z.number().min(0)),
});

export const createNewBatchValidation = (
  values: object,
): Result<CreateNewBatchValidationSchema, z.ZodError<object>> => {
  const validationResult = createNewBatchValidationSchema.safeParse(values);

  if (validationResult.success) {
    return { type: "ok", data: validationResult.data };
  } else {
    return { type: "err", data: validationResult.error };
  }
};
