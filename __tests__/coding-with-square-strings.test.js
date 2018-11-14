import fs from 'fs';
import { code, decode } from '../src/squared-strings/coding-with-square-strings';

const data1Sol = '\vctg?.nadr d gdbW\n\v,i    lnis tl eh\n\v mtIAakietboaara\n\veeo nnigsoe st?t\n\vd wsddnh lfls   \n\vgaaa  gtfeoeehWd\n'
  + '\vytrwbI .o rasiho\n\v, d e i rtev,se \n\v t hflnW h e  ny\n\vfhmioo emot Is o\n\voeemrvt eshh tIu\n\vr   eehw eaiwr  \n'
  + '\veptc deea tmaelr\n\viihot  rtc?.naoe\n\vgcamhhre h  tkom\n\vnntiaia meHAeyke\n\v.i ntmiwirend em';

const data2Sol = 'fa  h ttrheI ilS\nitifakw   s`irdo\nc cotnihftivce m\neAereocaihree,we\n.n   wedroe . i \n\vdIdT , es t Sls\n\v seoe t.eIaFola\n'
  + '\vw s nIo   srm y\n\voatso  Bwhtoee \n\vulrautpuhoem nt\n\vlsuyghetold sdh\n\vdoc hir  d wa e\n\v  tt niif ohyi \n\vsgihoksfawfa nw\n'
  + '\vuroaf h vi ti o\n\vfent I iotd nfr';

test('Code 0', () => {
  expect(code('')).toBe('');
});

test('Code 1', () => {
  const data1 = fs.readFileSync('./__tests__/__fixtures__/square-str.data1.txt', 'utf-8');
  // const data1Sol = fs.readFileSync('./__tests__/__fixtures__/square-str.solution1.txt', 'utf-8');
  expect(code(data1)).toEqual(data1Sol);
});

test('Code 2', () => {
  const data2 = fs.readFileSync('./__tests__/__fixtures__/square-str.data2.txt', 'utf-8');
  // const data2Sol = fs.readFileSync('./__tests__/__fixtures__/square-str.solution2.txt', 'utf-8');
  expect(code(data2)).toBe(data2Sol);
});

test('Decode 0', () => {
  expect(decode('')).toBe('');
});

test('Decode 1', () => {
  const data2 = fs.readFileSync('./__tests__/__fixtures__/square-str.data2.txt', 'utf-8');
  // const data2Sol = fs.readFileSync('./__tests__/__fixtures__/square-str.solution2.txt', 'utf-8');
  expect(decode(data2Sol)).toBe(data2);
});
