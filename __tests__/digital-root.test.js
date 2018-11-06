import digitalRoot from '../src/digital-root';

test('Sum of digitals 1', () => {
  expect(digitalRoot(0)).toBe(0);
});

test('Sum of digitals 2', () => {
  expect(digitalRoot(15)).toBe(6);
});

test('Sum of digitals 3', () => {
  expect(digitalRoot(7833)).toBe(3);
});

test('Sum of digitals 4', () => {
  expect(digitalRoot(555555)).toBe(3);
});
