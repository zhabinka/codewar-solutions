// https://www.codewars.com/kata/weight-for-weight/

import sumDigit from './sum-digital';

const sortWeight = (a, b) => (sumDigit(a) - sumDigit(b) || a.localeCompare(b));

const orderWeight = str => str
  .split(' ')
  .sort(sortWeight)
  .join(' ');

export default orderWeight;
