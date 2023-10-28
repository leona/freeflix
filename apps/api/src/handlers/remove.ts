import torrent from "../models/torrent";

export const remove = async (ctx) => {
  const { hash } = await ctx.req.json();
  await torrent.remove(hash);

  return ctx.json({
    message: "Removed from queue",
  });
};

export const removeTitle = async (ctx) => {
  const { title } = await ctx.req.json();
  await torrent.removeByTitle(title);

  return ctx.json({
    message: "Removed from queue",
  });
};
