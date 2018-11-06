import rgbToHex from '../src/convert-a-hex-string-to-rgb';

test('hex to rgb 0', () => {
  const rgb0 = { r: 0, g: 0, b: 0 };
  expect(rgbToHex('#000000')).toEqual(rgb0);
});

test('hex to rgb 1', () => {
  const rgb1 = { r: 255, g: 153, b: 51 };
  expect(rgbToHex('#FF9933')).toEqual(rgb1);
});

test('hex to rgb 2', () => {
  const rgb2 = { r: 221, g: 255, b: 211 };
  expect(rgbToHex('#DDFFD3')).toEqual(rgb2);
});

test('hex to rgb 3', () => {
  const rgb3 = { r: 187, g: 223, b: 210 };
  expect(rgbToHex('#bbdfd2')).toEqual(rgb3);
});
/*
test('hex to rgb shorthand form', () => {
  const rgb4 = { r: 204, g: 204, b: 204 };
  expect(rgbToHex('#CCC')).toEqual(rgb4);
});
*/
