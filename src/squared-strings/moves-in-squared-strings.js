// https://www.codewars.com/kata/moves-in-squared-strings-i/

const vertMirror = arr => arr.map(el => [...el].reverse().join(''));
const horMirror = arr => arr.reverse();
const selfie = arr => arr.map(el => el.padEnd(el.length * 2, '.'));

const rot = arr => horMirror(vertMirror(arr));

const selfieAndRot = (arr) => {
  const selfieArr = selfie(arr);
  return selfieArr.concat(rot(selfieArr));
};

const leftDiagMirror = (arr) => {
  const iter = (data, i) => {
    if (i === data.length) {
      return [];
    }
    const sideSquare = data.reduce((acc, el) => acc.concat(el[i]), '');

    return [sideSquare].concat(iter(data, i + 1));
  };

  return iter(arr, 0);
};

const rightDiagMirror = arr => leftDiagMirror(rot(arr));

const rot90Clock = arr => vertMirror(leftDiagMirror(arr));

const selfieAndDiag = (arr) => {
  const diagonaleArr = leftDiagMirror(arr);
  return arr.map((el, i) => [el].concat(diagonaleArr[i]).join('|'));
};

const oper = (f, data) => f(data.split('\n')).join('\n');

export {
  oper,
  vertMirror,
  horMirror,
  rot,
  selfieAndRot,
  leftDiagMirror,
  rot90Clock,
  selfieAndDiag,
  rightDiagMirror,
  // rot90Counter,
  // selfieDiag2Counterclock,
};
