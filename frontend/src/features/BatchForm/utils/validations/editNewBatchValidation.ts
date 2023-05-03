import z from "zod";
import type { components } from "~/swagger/schema/schemas/batch";
import { Result } from "~/types/Result";

type EditBatchValidationSchema = components["schemas"]["Batch"];
const editBatchValidationSchema: z.ZodSchema<EditBatchValidationSchema, z.ZodTypeDef, object> =
  z.object({
    id: z.number(),
    batch_name: z.string().min(1),
    server_name: z.string().min(1),
    cron_setting: z.string().min(1),
    initial_date: z.date().transform((date) => {
      return date.toISOString();
    }),
    time_limit: z.preprocess(Number, z.number().min(0)),
  });

export const editBatchValidation = (
  values: object,
): Result<EditBatchValidationSchema, z.ZodError<object>> => {
  const validationResult = editBatchValidationSchema.safeParse(values);

  if (validationResult.success) {
    return { type: "ok", data: validationResult.data };
  } else {
    return { type: "err", data: validationResult.error };
  }
};
