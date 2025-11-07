// https://www.codewars.com/kata/ipv4-to-int32/

const binaryIp = ip => ip
  .split('.')
  .map(el => (Number(el)).toString(2).padStart(8, 0))
  .join('');

const ipToInt32 = ip => parseInt(binaryIp(ip), 2);

export default ipToInt32;
