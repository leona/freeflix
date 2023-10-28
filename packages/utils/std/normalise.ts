export const floatToPercent = (input) => {
  return `${parseInt(input * 100)}%`;
};

export const bytes = (speed, measurements = ["KiB/S", "MiB/s", "GiB/s"]) => {
  const [kb, mb, gb] = measurements;
  let kibSpeed = speed / 1024;

  if (kibSpeed < 1024) {
    return `${kibSpeed.toFixed(2)} ${kb}`;
  }

  let mibSpeed = kibSpeed / 1024;

  if (mibSpeed < 1024) {
    return `${mibSpeed.toFixed(2)} ${mb}`;
  }

  let gibSpeed = mibSpeed / 1024;
  return `${gibSpeed.toFixed(2)} ${gb}`;
};
