// https://www.codewars.com/kata/int32-to-ipv4

const int32ToIp = (code) => {
  const ipStr = code.toString(2);

  return '0'
    .repeat(32 - ipStr.length)
    .concat(ipStr)
    .match(/.{1,8}/g)
    .map(el => parseInt(el, 2))
    .join('.');
};

export default int32ToIp;
