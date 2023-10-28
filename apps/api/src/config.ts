import { XMLParser } from "fast-xml-parser";
import fs from "fs";

try {
  const prowlarrConfigXML = fs.readFileSync(
    "/app/config/prowlarr/config.xml",
    "utf8"
  );

  const parser = new XMLParser();
  var prowlarrConfig = parser.parse(prowlarrConfigXML).Config;
} catch (err) {
  console.error("Failed to get prowlarr config, restarting.");
  await new Promise((r) => setTimeout(r, 5000));
  process.exit();
}

export const prowlarr = prowlarrConfig;
