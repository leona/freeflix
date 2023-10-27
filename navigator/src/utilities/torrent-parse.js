function parseTorrentFilename(filename) {
  let parsedData = {};
  let filenameNorm = filename.replace(/[\.]/g, " ");

  // extracting tv/movie name
  let nameExp =
    /^(.*?)(s\d+e\d+|[\d]{4}|complete|s\d+)?(.*)(x264|x265)?(.*)(hdtv|webrip|web-dl|bluray)?(.*?)(-.*?)?(mkv|mp4|avi)?$/i;
  let extractName = nameExp.exec(filenameNorm);

  parsedData.name = extractName[1].trim();

  // extracting season & episode
  let seasonEpisodeExp = /(s\d+e\d+|s\d+)?/i;
  let extractSeasonEpisode = seasonEpisodeExp.exec(filenameNorm);

  if (extractSeasonEpisode[0]) {
    let seasonEpisodeSplit = extractSeasonEpisode[0].split("e");
    parsedData.season = parseInt(seasonEpisodeSplit[0].replace("s", ""));
    parsedData.episode = parseInt(seasonEpisodeSplit[1]);
  } else {
    parsedData.season = null;
    parsedData.episode = null;
  }

  // extracting quality
  let qualityExp = /(hdtv|webrip|bluray|web-dl)/i;
  let extractQuality = qualityExp.exec(filenameNorm);

  parsedData.quality = extractQuality ? extractQuality[0].toUpperCase() : null;

  // extracting group
  let groupExp = /(-.*?)($|\[.*\])*/g;
  let extractGroup = groupExp.exec(filenameNorm);

  parsedData.group = extractGroup
    ? extractGroup[1].replace("-", "").trim().toUpperCase()
    : null;

  return parsedData;
}
