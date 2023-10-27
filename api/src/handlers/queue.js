import torrent from "../models/torrent";

export const queue = async (ctx) => {
  const { url } = await ctx.req.json();
  await torrent.add(url);

  return ctx.json({
    message: "Added to queue",
  });
};
