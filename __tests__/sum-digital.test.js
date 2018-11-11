import sumDigital from '../src/sum-digital';

test('Sum of digitals 1', () => {
  expect(sumDigital(0)).toBe(0);
});

test('Sum of digitals 2', () => {
  expect(sumDigital(15)).toBe(6);
});

test('Sum of digitals 3', () => {
  expect(sumDigital(7833)).toBe(21);
});

test('Sum of digitals 4', () => {
  expect(sumDigital(555555)).toBe(30);
});
