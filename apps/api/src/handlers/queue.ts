import torrent from "../models/torrent";

export const queue = async (ctx) => {
  const { url } = await ctx.req.json();

  try {
    await torrent.add(url);
  } catch (e) {
    return ctx.json({ error: e.message }, 500);
  }

  return ctx.json({
    message: "Added to queue",
  });
};
