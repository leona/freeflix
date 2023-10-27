import prowlarr from "../models/prowlarr.js";

export const search = async (ctx) => {
  const query = ctx.req.query("query");
  const results = await prowlarr.search({ query });
  return ctx.json(results);
};
