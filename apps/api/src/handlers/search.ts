import torrent from "../models/torrent.js";

export const search = async (ctx) => {
  const query = ctx.req.query("query");
  const results = await torrent.search(query);
  return ctx.json(results);
};
