import jellyfin from "../models/jellyfin.js";

export const watch = async (ctx) => {
  const query = ctx.req.query("query");
  const results = await jellyfin.search({ query });

  if (!results?.SearchHints?.length) {
    return ctx.json({
      url: `https://${process.env.PUBLIC_URL}/jellyfin/web/index.html#!/home.html`,
    });
  }

  return ctx.json({
    url: `https://${process.env.PUBLIC_URL}/jellyfin/web/index.html#!/details?id=${results.SearchHints[0].Id}`,
  });
};
