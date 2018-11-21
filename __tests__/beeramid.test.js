import beeramid from '../src/beeramid';

test('Beeramid 0/1', () => {
  expect(beeramid(3, 4)).toBe(0);
});

test('Beeramid 0/2', () => {
  expect(beeramid(0, 4)).toBe(0);
});

test('Beeramid 0/3', () => {
  expect(beeramid(-1, 4)).toBe(0);
});

test('Beeramid 1', () => {
  expect(beeramid(9, 2)).toBe(1);
});

test('Beeramid 2', () => {
  expect(beeramid(10, 2)).toBe(2);
});

test('Beeramid 3', () => {
  expect(beeramid(11, 2)).toBe(2);
});

test('Beeramid 4', () => {
  expect(beeramid(21, 1.5)).toBe(3);
});

test('Beeramid 5', () => {
  expect(beeramid(454, 5)).toBe(5);
});

test('Beeramid 6', () => {
  expect(beeramid(455, 5)).toBe(6);
});

test('Beeramid 7', () => {
  expect(beeramid(4, 4)).toBe(1);
});
