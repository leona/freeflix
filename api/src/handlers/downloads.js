import torrent from "../models/torrent";

export const downloads = async (ctx) => {
  const data = await torrent.downloads();
  return ctx.json(data);
};
