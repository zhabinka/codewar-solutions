import strInc from '../src/string-incrementer';

test('String inc 0', () => {
  expect(strInc('')).toBe('1');
});

test('String inc 1', () => {
  expect(strInc('foo')).toBe('foo1');
});

test('String inc 2', () => {
  expect(strInc('foobar23')).toBe('foobar24');
});

test('String inc 3', () => {
  expect(strInc('foo0042')).toBe('foo0043');
});

test('String inc 4', () => {
  expect(strInc('foo9')).toBe('foo10');
});

test('String inc 5', () => {
  expect(strInc('foo099')).toBe('foo100');
});

test('String inc 6', () => {
  expect(strInc('foo99')).toBe('foo100');
});
