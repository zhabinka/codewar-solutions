// https://www.codewars.com/kata/weight-for-weight/

import sumDigit from './sum-digital';

const sortWeight = (a, b) => {
  if (a.weight === b.weight) {
    return (a.num).localeCompare(b.num);
  }

  return a.weight - b.weight;
};

const orderWeight = str => str
  .split(' ')
  .map(el => ({ num: el, weight: sumDigit(el) }))
  .sort(sortWeight)
  .reduce((acc, el) => acc.concat(el.num), [])
  .join(' ');

export default orderWeight;
