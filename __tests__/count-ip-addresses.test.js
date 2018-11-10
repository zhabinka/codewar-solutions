import ipsBetween from '../src/count-ip-addresses';

test('Count IP adresses 1', () => {
  expect(ipsBetween('10.0.0.0', '10.0.0.50')).toBe(50);
});

test('Count IP adresses 2', () => {
  expect(ipsBetween('10.0.0.0', '10.0.1.0')).toBe(256);
});

test('Count IP adresses 3', () => {
  expect(ipsBetween('20.0.0.10', '20.0.1.0')).toBe(246);
});
