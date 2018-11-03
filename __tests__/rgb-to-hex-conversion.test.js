import rgb from '../src/rgb-to-hex-conversion';

test('rgb to hex 1', () => {
  expect(rgb(255, 255, 255)).toBe('FFFFFF');
});

test('rgb to hex 2', () => {
  expect(rgb(255, 255, 300)).toBe('FFFFFF');
});

test('rgb to hex 3', () => {
  expect(rgb(0, 0, 0)).toBe('000000');
});

test('rgb to hex 4', () => {
  expect(rgb(148, 0, 211)).toBe('9400D3');
});
