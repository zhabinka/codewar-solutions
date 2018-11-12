import {
  oper,
  vertMirror,
  horMirror,
  rot,
  selfieAndRot,
} from '../src/squared-strings/moves-in-squared-strings';

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

test('Rot Mirror 1', () => {
  const mirror5 = 'BClf\nhkXz\nMfoc\nvKkr';
  expect(oper(rot, 'rkKv\ncofM\nzXkh\nflCB')).toBe(mirror5);
});

test('Rot Mirror 2', () => {
  const mirror6 = 'LyfUVC\nGUFjWk\nABpfre\nJMmPrD\ntcVYqC\nooujif';
  expect(oper(rot, 'fijuoo\nCqYVct\nDrPmMJ\nerfpBA\nkWjFUG\nCVUfyL')).toBe(mirror6);
});

test('Selfie and Rot Mirror 1', () => {
  const mirror7 = 'xZBV....\njsbS....\nJcpN....\nfVnP....\n....PnVf\n....NpcJ\n....Sbsj\n....VBZx';
  expect(oper(selfieAndRot, 'xZBV\njsbS\nJcpN\nfVnP')).toBe(mirror7);
});

test('Selfie and Rot Mirror 2', () => {
  const mirror8 = 'uLcq....\nJkuL....\nYirX....\nnwMB....\n....BMwn\n....XriY\n....LukJ\n....qcLu';
  expect(oper(selfieAndRot, 'uLcq\nJkuL\nYirX\nnwMB')).toBe(mirror8);
});

/*
test('Rot Mirror 1', () => {
  const mirror5 = '';
  expect(oper(rot, '')).toBe(mirror5);
});
*/
