import { oper, vertMirror, horMirror } from '../src/squared-strings/moves-in-squared-strings';

test('Vertical Mirror 1', () => {
  const mirror1 = 'QHdgSh\noaMDnH\nXxNNlC\nHxxvRi\nAvVTqb\nuRySvw';
  expect(oper(vertMirror, 'hSgdHQ\nHnDMao\nClNNxX\niRvxxH\nbqTVvA\nwvSyRu')).toBe(mirror1);
});

test('Vertical Mirror 2', () => {
  const mirror2 = 'EWTOzI\nMCebkk\nMxZzuW\nwJddDv\nFHyJij\nxSfHVP';
  expect(oper(vertMirror, 'IzOTWE\nkkbeCM\nWuzZxM\nvDddJw\njiJyHF\nPVHfSx')).toBe(mirror2);
});

test('Horizontal Mirror 1', () => {
  const mirror3 = 'yeCt\nCSbg\nJVhv\nlVHt';
  expect(oper(horMirror, 'lVHt\nJVhv\nCSbg\nyeCt')).toBe(mirror3);
});

test('Horizontal Mirror 2', () => {
  const mirror4 = 'cEYz\nLPKo\ndbrZ\nnjMK';
  expect(oper(horMirror, 'njMK\ndbrZ\nLPKo\ncEYz')).toBe(mirror4);
});
