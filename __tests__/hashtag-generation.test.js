import generateHashtag from '../src/hashtag-generator';

test('Hashtag generator 1', () => {
  expect(generateHashtag(' Hello there thanks for trying my Kata')).toBe('#HelloThereThanksForTryingMyKata');
});

test('Hashtag generator 2', () => {
  expect(generateHashtag('    Hello     World   ')).toBe('#HelloWorld');
});

test('Hashtag generator 3', () => {
  expect(generateHashtag('')).toBe(false);
});
