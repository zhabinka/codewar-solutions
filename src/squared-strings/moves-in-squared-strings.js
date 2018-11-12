// https://www.codewars.com/kata/moves-in-squared-strings-i/

const vertMirror = arr => arr.map(el => [...el].reverse().join(''));
const horMirror = arr => arr.reverse();

const oper = (f, data) => f(data.split('\n')).join('\n');

export { oper, vertMirror, horMirror };
