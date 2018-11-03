// https://www.codewars.com/kata/rgb-to-hex-conversion/train/javascript

const checkValue = (x) => {
  switch (true) {
    case (x < 0):
      return 0;
    case (x > 255):
      return 255;
    default:
      return x;
  }
};

const rgb = (...values) => values
  .map(el => checkValue(el)
    .toString(16)
    .replace(0, '00'))
  .join('')
  .toUpperCase();

export default rgb;
