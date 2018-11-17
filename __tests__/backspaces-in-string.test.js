import processString from '../src/backspaces-in-string';

test('Backspaces in string 0', () => {
  expect(processString('')).toBe('');
});

test('Backspaces in string 1', () => {
  expect(processString('abc#d##c')).toBe('ac');
});

test('Backspaces in string 2', () => {
  expect(processString('abc####d##c#')).toBe('');
});

test('Backspaces in string 3', () => {
  expect(processString('###########')).toBe('');
});

test('Backspaces in string 4', () => {
  expect(processString('831####jns###s#cas/*####-5##s##6+yqw87e##hfklsd-=-28##fds##')).toBe('6+yqw8hfklsd-=-f');
});

test('Backspaces in string 5', () => {
  expect(processString('fjnwui#NI#(*N#ION#Onfjen################Io4f')).toBe('Io4f');
});
