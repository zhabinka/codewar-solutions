// https://www.codewars.com/kata/sum-of-digits-slash-digital-root

import sumDigit from './sum-digital';

const digitalRoot = num => (num < 10
  ? num
  : digitalRoot(sumDigit(num)));

export default digitalRoot;
