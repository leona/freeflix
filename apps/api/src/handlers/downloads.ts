import torrent from "../models/torrent";
import fs from "fs";

export const downloads = async (ctx) => {
  const data = await torrent.downloads();
  return ctx.json(data);
};

export const download = async (ctx) => {
  const query = ctx.req.query("query");

  if (query.includes("..") || query.includes("/")) {
    return ctx.json(
      {
        error: "invalid path",
      },
      400
    );
  }

  const filepath = findLargestFileInPath("/app/data/media/" + query);
  const filename = filepath.split("/").pop();
  const file = Bun.file(filepath);
  const response = new Response(file);

  response.headers.set(
    "Content-Disposition",
    `attachment; filename=${filename}`
  );

  return response;
};

const findLargestFileInPath = (path): string => {
  if (!fs.lstatSync(path).isDirectory()) {
    return path;
  }

  const files = fs.readdirSync(path);

  const largest = files
    .map((file) => {
      const stat = fs.statSync(`${path}/${file}`);
      return {
        file,
        size: stat.size,
      };
    })
    .sort((a, b) => b.size - a.size)[0];

  return `${path}/${largest.file}`;
};
