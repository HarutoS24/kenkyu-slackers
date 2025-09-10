import { getAxios } from "@/util/axios";
import z from "zod";

const getFeedbackFromGPTRequestSchema = z.object({
  text: z.string(),
  release_type_id: z.string(),
  important_aspects: z.array(z.string()),
});
type GetFeedbackFromGPTRequestSchema = z.infer<typeof getFeedbackFromGPTRequestSchema>;

const getFeedbackFromGPTResponseSchema = z.object({
  advice: z.string(),
  improved_press: z.string(),
});

export const getFeedbackFromGPT = async (req: GetFeedbackFromGPTRequestSchema) => {
  getFeedbackFromGPTRequestSchema.parse(req);

  const res = await getAxios().post("/get_feedback_from_GPT", req);
  if (res.status === 200) {
    const result = getFeedbackFromGPTResponseSchema.parse(res.data);
    return result;
  }
  else {
    throw new Error(`The API response was unexpected (status ${res.status}).`);
  }
}
