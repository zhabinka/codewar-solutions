// https://www.codewars.com/kata/sum-of-digits-slash-digital-root

const digital_root = (num) => {
  return num < 10
    ? num
    : digital_root([...String(num)].reduce((acc, n) => acc + Number(n), 0));
};