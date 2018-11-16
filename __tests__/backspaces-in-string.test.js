import processString from '../src/backspaces-in-string';

test('Backspaces in string 0', () => {
  expect(processString('')).toBe('');
});

test('Backspaces in string 1', () => {
  expect(processString('abc#d##c')).toBe('ac');
});

test('Backspaces in string 2', () => {
  expect(processString('abc####d##c#')).toBe('');
});

test('Backspaces in string 3', () => {
  expect(processString('###########')).toBe('');
});
