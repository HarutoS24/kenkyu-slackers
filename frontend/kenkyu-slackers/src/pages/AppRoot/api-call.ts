import { getAxios } from "@/util/axios";
import z from "zod";

const getIndustryIdsResponseSchema = z.record(z.string(), z.string());

export const getIndustryIds = async () => {
  const res = await getAxios().get("/industry_ids");
  if (res.status === 200) {
    const result = getIndustryIdsResponseSchema.parse(res.data);
    return result;
  }
  else {
    throw new Error(`The API response was unexpected (status ${res.status}).`);
  }
}

const getFeedbackFromGPTResponseSchema = z.object({
  Advice: z.string(),
  improved_press: z.string(),
});

export const getFeedbackFromGPT = async () => {
  const res = await getAxios().post("/get_feedback_from_GPT", {});
  if (res.status === 200) {
    const result = getFeedbackFromGPTResponseSchema.parse(res.data);
    return result;
  }
  else {
    throw new Error(`The API response was unexpected (status ${res.status}).`);
  }
}
