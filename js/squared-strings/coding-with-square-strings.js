// https://www.codewars.com/kata/coding-with-squared-strings/

import { rot90Clock, rot90Counter } from './moves-in-squared-strings';

const code = (str) => {
  if (str.length === 0) {
    return '';
  }
  const sizeOfSquare = Math.ceil(Math.sqrt(str.length));
  const re = new RegExp(`.{${sizeOfSquare}}`, 'g');
  const arr = str
    .padEnd(sizeOfSquare ** 2, '\v')
    .match(re);

  return rot90Clock(arr).join('\n');
};

const decode = (str) => {
  if (str.length === 0) {
    return '';
  }

  return rot90Counter(str.split('\n'))
    .join('')
    .trim();
};

export { code, decode };
