// https://www.codewars.com/kata/moves-in-squared-strings-i/

const vertMirror = arr => arr.map(el => [...el].reverse().join(''));
const horMirror = arr => arr.reverse();
const selfie = arr => arr.map(el => el.padEnd(el.length * 2, '.'));

const rot = arr => horMirror(vertMirror(arr));

const selfieAndRot = (arr) => {
  const selfieArr = selfie(arr);
  return selfieArr.concat(rot(selfieArr));
};

const diagMirror = (arr) => {
  const iter = (data, i) => {
    if (i === data.length) {
      return [];
    }
    const sideSquare = data.reduce((acc, el) => acc.concat(el[i]), '');

    return [sideSquare].concat(iter(data, i + 1));
  };

  return iter(arr, 0);
};

const rot90Clock = arr => vertMirror(diagMirror(arr));

const selfieAndDiag = (arr) => {
  const diagonaleArr = diagMirror(arr);
  return arr.map((el, i) => [el].concat(diagonaleArr[i]).join('|'));
};

const oper = (f, data) => f(data.split('\n')).join('\n');

export {
  oper,
  vertMirror,
  horMirror,
  rot,
  selfieAndRot,
  diagMirror,
  rot90Clock,
  selfieAndDiag,
};
