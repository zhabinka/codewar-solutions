// https://www.codewars.com/kata/moves-in-squared-strings-i/

const vertMirror = arr => arr.map(el => [...el].reverse().join(''));
const horMirror = arr => arr.reverse();
const selfie = arr => arr.map(el => el.padEnd(el.length * 2, '.'));

const rot = arr => horMirror(vertMirror(arr));

const selfieAndRot = (arr) => {
  const selfieArr = selfie(arr);
  return selfieArr.concat(rot(selfieArr));
};

const oper = (f, data) => f(data.split('\n')).join('\n');

export {
  oper,
  vertMirror,
  horMirror,
  rot,
  selfieAndRot,
};
