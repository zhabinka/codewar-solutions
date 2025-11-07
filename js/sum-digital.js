const sumDigital = num => num
  .toString()
  .split('')
  .reduce((acc, n) => acc + Number(n), 0);

export default sumDigital;
