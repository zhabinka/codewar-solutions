import findNum from '../src/unique-digit-sequence';

test('Unique digit sequence 1', () => {
  expect(findNum(1)).toBe(1);
});

test('Unique digit sequence 2', () => {
  expect(findNum(5)).toBe(5);
});

test('Unique digit sequence 3', () => {
  expect(findNum(11)).toBe(22);
});

test('Unique digit sequence 4', () => {
  expect(findNum(100)).toBe(103);
});

test('Unique digit sequence 5', () => {
  expect(findNum(500)).toBe(476);
});
