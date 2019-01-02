import * as math from '../src/math-issues';

test('Math.round(0.4)', () => {
  expect(math.round(0.4)).toEqual(0);
});

test('Math.round(0.5)', () => {
  expect(math.round(0.5)).toEqual(1);
});

test('Math.ceil(0.4)', () => {
  expect(math.ceil(0.4)).toEqual(1);
});

test('Math.ceil(0.5)', () => {
  expect(math.ceil(0.5)).toEqual(1);
});

test('Math.floor(0.4)', () => {
  expect(math.floor(0.4)).toEqual(0);
});

test('Math.floor(0.5)', () => {
  expect(math.floor(0.5)).toEqual(0);
});
