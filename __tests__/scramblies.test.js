import { scramble, scrambleLoop } from '../src/scramblies';

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

test('Scramble inc 6', () => {
  expect(scrambleLoop('soppwtbqbpszexjgcswawzdmgfdsbukivgiiwqccrucylasrsrrjqhgbplnnvkzfdumkepbywipphboqiegzeggtbbvkatwurkctfqvsiyqynenxahkdfepphgngsazmbwbpbphtedajwoogyzwqtcdrwsxxuatxtmhfnsxdiioyemevp', 'zwigaumywtgltlbrwanqxujvawphorynmtahlaxaxkavknzbxatfczobnwnwauhmiktgrqjnowuqhoyfluahnsfkmmprccp')).toBe(false);
});

test('Scramble inc 7', () => {
  expect(scrambleLoop('jscripts', 'javascript')).toBe(false);
});

test('Scramble inc 8', () => {
  expect(scrambleLoop('aabbcamaomsccdd', 'commas')).toBe(true);
});

test('Scramble loop 1', () => {
  expect(scrambleLoop('rkqodlw', 'world')).toBe(true);
});

test('Scramble loop 2', () => {
  expect(scrambleLoop('cedewaraaossoqqyt', 'codewars')).toBe(true);
});

test('Scramble loop 3', () => {
  expect(scrambleLoop('katas', 'steak')).toBe(false);
});

test('Scramble loop 4', () => {
  expect(scrambleLoop('scriptjava', 'javascript')).toBe(true);
});

test('Scramble loop 5', () => {
  expect(scrambleLoop('scriptingjava', 'javascript')).toBe(true);
});

test('Scramble loop 6', () => {
  expect(scrambleLoop('soppwtbqbpszexjgcswawzdmgfdsbukivgiiwqccrucylasrsrrjqhgbplnnvkzfdumkepbywipphboqiegzeggtbbvkatwurkctfqvsiyqynenxahkdfepphgngsazmbwbpbphtedajwoogyzwqtcdrwsxxuatxtmhfnsxdiioyemevp', 'zwigaumywtgltlbrwanqxujvawphorynmtahlaxaxkavknzbxatfczobnwnwauhmiktgrqjnowuqhoyfluahnsfkmmprccp')).toBe(false);
});

test('Scramble loop 7', () => {
  expect(scrambleLoop('jscripts', 'javascript')).toBe(false);
});

test('Scramble loop 8', () => {
  expect(scrambleLoop('aabbcamaomsccdd', 'commas')).toBe(true);
});
