// https://www.codewars.com/kata/math-issues/

const parseNumeric = num => num.toString().split('.');


const round = (number) => {
  const [int, decimal] = parseNumeric(number);

  return [...decimal].slice(0, 1) < 5 ? +int : +int + 1;
};

const ceil = (number) => {
  const [int, decimal] = parseNumeric(number);

  return +decimal > 0 ? +int + 1 : +int;
};

const floor = (number) => {
  const [int] = parseNumeric(number);

  return +int;
};

export { round, ceil, floor };
