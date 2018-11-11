import orderWeight from '../src/weight-for-weight';

test('Order weight 1', () => {
  const str1 = '100 180 90 56 65 74 68 86 99';
  expect(orderWeight('56 65 74 100 99 68 86 180 90')).toBe(str1);
});

test('Order weight 2', () => {
  const str2 = '2000 103 123 4444 99';
  expect(orderWeight('103 123 4444 99 2000')).toBe(str2);
});

test('Order weight 3', () => {
  const str3 = '11 11 2000 10003 22 123 1234000 44444444 9999';
  expect(orderWeight('2000 22 10003 1234000 44444444 9999 11 11 123')).toBe(str3);
});

test('Order weight 4', () => {
  const str4 = '112 140 114 24 6 115 9 65 6520 68 313611 97 198 322291 412462 16385 342437 73905 415268 77587 368488';
  expect(orderWeight('73905 115 322291 24 6520 97 412462 68 313611 65 77587 114 16385 198 415268 9 368488 112 342437 140 6')).toBe(str4);
});
