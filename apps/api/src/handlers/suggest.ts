import imdb from "../models/imdb";

export const suggest = async (ctx) => {
  const query = ctx.req.query("query");

  try {
    const results = await imdb.suggest({ query });
    return ctx.json(results);
  } catch (e) {
    return ctx.json({ error: e.message }, 500);
  }
};
