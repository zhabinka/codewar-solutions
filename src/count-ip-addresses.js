// https://www.codewars.com/kata/count-ip-addresses/

import ipToInt32 from './ipv4-to-int32';

const ipsBetween = (start, end) => ipToInt32(end) - ipToInt32(start);

export default ipsBetween;
