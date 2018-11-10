import whoIsNext from '../src/double-cola';

test('Who will drink N-th can 1', () => {
  expect(whoIsNext(['Sheldon', 'Leonard', 'Penny', 'Rajesh', 'Howard'], 1)).toBe('Sheldon');
});

test('Who will drink N-th can 2', () => {
  expect(whoIsNext(['Sheldon', 'Leonard', 'Penny', 'Rajesh', 'Howard'], 52)).toBe('Penny');
});

test('Who will drink N-th can 3', () => {
  expect(whoIsNext(['Sheldon', 'Leonard', 'Penny', 'Rajesh', 'Howard'], 7230702951)).toBe('Leonard');
});
