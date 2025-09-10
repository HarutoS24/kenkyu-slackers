import { defineStore } from 'pinia';

export const useReviewContentStore = defineStore("markdown-content", {
  state: (): {
    markdownContent: string,
    industryIds: string[],
    importantAspects: string[],
  } => ({
    markdownContent: "",
    industryIds: [],
    importantAspects: [],
  }),
  getters: {
    singleIndustryId: (state) => state.industryIds[0]
  }
});
