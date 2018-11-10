import ipToInt32 from '../src/ipv4-to-int32';

test('IPv4 to integer 0', () => {
  expect(ipToInt32('0.0.0.0')).toBe(0);
});

test('IPv4 to integer 1', () => {
  expect(ipToInt32('0.0.0.32')).toBe(32);
});

test('IPv4 to integer 2', () => {
  expect(ipToInt32('128.32.10.1')).toBe(2149583361);
});

test('IPv4 to integer 3', () => {
  expect(ipToInt32('128.114.17.104')).toBe(2154959208);
});
