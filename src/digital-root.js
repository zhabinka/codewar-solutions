// https://www.codewars.com/kata/sum-of-digits-slash-digital-root

const digitalRoot = num => (num < 10
  ? num
  : digitalRoot(num
    .toString()
    .split('')
    .reduce((acc, n) => acc + Number(n), 0)));

export default digitalRoot;
