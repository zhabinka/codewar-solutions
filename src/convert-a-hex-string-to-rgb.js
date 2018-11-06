// https://www.codewars.com/kata/rgb-to-hex-conversion

const rgbToHex = (code) => {
  const [r, g, b] = code
    .replace(/^#/, '')
    .match(/.{2}/g)
    .map(el => parseInt(el, 16));

  return { r, g, b };
};

export default rgbToHex;
