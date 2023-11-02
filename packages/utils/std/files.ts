import fs from "fs";

export const findLargestFileInPath = (path): string => {
  console.log("finding largest file in path:", path);

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
