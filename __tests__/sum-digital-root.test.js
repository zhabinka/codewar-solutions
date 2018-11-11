import digitalRoot from '../src/sum-digital-root';

test('Sum of digitals root 1', () => {
  expect(digitalRoot(0)).toBe(0);
});

test('Sum of digitals root 2', () => {
  expect(digitalRoot(15)).toBe(6);
});

test('Sum of digitals root 3', () => {
  expect(digitalRoot(7833)).toBe(3);
});

test('Sum of digitals root 4', () => {
  expect(digitalRoot(555555)).toBe(3);
});
