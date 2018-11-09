import scramble from '../src/scramblies';

test('Scramble 1', () => {
  expect(scramble('rkqodlw', 'world')).toBe(true);
});

test('Scramble 2', () => {
  expect(scramble('cedewaraaossoqqyt', 'codewars')).toBe(true);
});

test('Scramble 3', () => {
  expect(scramble('katas', 'steak')).toBe(false);
});

test('Scramble 4', () => {
  expect(scramble('scriptjava', 'javascript')).toBe(true);
});

test('Scramble 5', () => {
  expect(scramble('scriptingjava', 'javascript')).toBe(true);
});

test('String inc 6', () => {
  expect(scramble('soppwtbqbpszexjgcswawzdmgfdsbukivgiiwqccrucylasrsrrjqhgbplnnvkzfdumkepbywipphboqiegzeggtbbvkatwurkctfqvsiyqynenxahkdfepphgngsazmbwbpbphtedajwoogyzwqtcdrwsxxuatxtmhfnsxdiioyemevp', 'zwigaumywtgltlbrwanqxujvawphorynmtahlaxaxkavknzbxatfczobnwnwauhmiktgrqjnowuqhoyfluahnsfkmmprccp')).toBe(false);
});

test('String inc 7', () => {
  expect(scramble('jscripts', 'javascript')).toBe(false);
});

test('String inc 8', () => {
  expect(scramble('aabbcamaomsccdd', 'commas')).toBe(true);
});
