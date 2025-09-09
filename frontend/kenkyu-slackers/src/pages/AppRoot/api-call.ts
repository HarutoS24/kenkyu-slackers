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
