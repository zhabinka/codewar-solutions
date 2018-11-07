import int32ToIp from '../src/int32-to-ipv4';

test('Integer to IPv4 0', () => {
  expect(int32ToIp(0)).toBe('0.0.0.0');
});

test('Integer to IPv4 1', () => {
  expect(int32ToIp(32)).toBe('0.0.0.32');
});

test('Integer to IPv4 2', () => {
  expect(int32ToIp(2149583361)).toBe('128.32.10.1');
});

test('Integer to IPv4 3', () => {
  expect(int32ToIp(2154959208)).toBe('128.114.17.104');
});
