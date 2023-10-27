const suggest = async ({ query }) => {
  const response = await fetch(
    `https://v3.sg.media-imdb.com/suggestion/x/${query}.json`
  );

  if (!response.ok) {
    throw new Error("Failed to get suggestions");
  }

  const data = await response.json();

  return data?.d.map((item) => ({
    id: item.id,
    title: item.l,
    year: item.y,
    image: item.i?.imageUrl,
    type: item.qid,
  }));
};

export default {
  suggest,
};
