import torrent from "../models/torrent";
import { sign, verify } from "hono/jwt";
import { findLargestFileInPath } from "@packages/utils/std/files";

export const downloads = async (ctx) => {
  const data = await torrent.downloads();
  return ctx.json(data);
};

export const generateDownload = async (ctx) => {
  const query = ctx.req.query("name");
  const data = await torrent.downloads();

  if (!data.complete.find((item) => item.name === query)) {
    return ctx.json(
      {
        error: "not found",
      },
      400
    );
  }

  const expiry = new Date();
  expiry.setHours(expiry.getHours() + 4);

  const token = await sign(
    {
      query,
      expiry: expiry.getTime(),
    },
    process.env.JWT_SECRET
  );

  return ctx.json({
    token,
  });
};

export const download = async (ctx) => {
  const query = ctx.req.query("token");

  if (!query) {
    return ctx.json(
      {
        error: "no token",
      },
      400
    );
  }

  try {
    var decoded = await verify(query, process.env.JWT_SECRET);
  } catch (e) {
    return ctx.json(
      {
        error: "invalid token",
      },
      400
    );
  }

  if (decoded.expiry < new Date().getTime()) {
    return ctx.json(
      {
        error: "expired token",
      },
      400
    );
  }

  const filepath = findLargestFileInPath(`/media/${decoded.query}`);
  const filename = filepath.split("/").pop();
  const file = Bun.file(filepath);
  const response = new Response(file);

  response.headers.set(
    "Content-Disposition",
    `attachment; filename=${filename}`
  );

  return response;
};
