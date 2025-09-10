import { getAxios } from "@/util/axios";
import z from "zod";

const getIndustryIdsResponseSchema = z.record(z.string(), z.string());

export const getIndustryIds = async () => {
  const res = await getAxios().get("/release_type");
  if (res.status === 200) {
    const result = getIndustryIdsResponseSchema.parse(res.data);
    return result;
  }
  else {
    throw new Error(`The API response was unexpected (status ${res.status}).`);
  }
}

// eslint-disable-next-line
const getAspectsResponseSchema = z.record(z.string(), z.string());

export const getAspects = async () => {
  // TODO: コピペなので直す
  const res = await getAxios().get("/aspect");
  if (res.status === 200) {
    const result = getIndustryIdsResponseSchema.parse(res.data);
    return result;
  }
  else {
    throw new Error(`The API response was unexpected (status ${res.status}).`);
  }
}
