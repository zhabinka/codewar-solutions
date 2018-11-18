import findIndexOf from '../src/number-in-array';

test('Find index of number 0', () => {
  expect(findIndexOf([5, 6, 1, 2, 3, 4], 3)).toBe(4);
});

test('Find index of number 1', () => {
  expect(findIndexOf([9, 11, 12, 1, 3, 5, 7], 5)).toBe(5);
});

test('Find index of number 2', () => {
  expect(findIndexOf([9, 11, 12, 1, 3, 5, 7], 2)).toBe(-1);
});
