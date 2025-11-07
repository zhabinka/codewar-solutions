// https://www.codewars.com/kata/int32-to-ipv4

const int32ToIp = code => code
  .toString(2)
  .padStart(32, 0)
  .match(/.{8}/g)
  .map(el => parseInt(el, 2))
  .join('.');

export default int32ToIp;
