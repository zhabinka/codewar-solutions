// https://www.codewars.com/kata/moves-in-squared-strings-i/

const vertMirror = str => str
  .split('\n')
  .map(el => [...el].reverse().join(''))
  .join('\n');

const horMirror = str => str
  .split('\n')
  .reverse()
  .join('\n');

const oper = (f, data) => f(data);

export { oper, vertMirror, horMirror };
