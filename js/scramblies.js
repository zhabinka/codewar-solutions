// https://www.codewars.com/kata/scramblies/

const scramble = (chars, str) => {
  if (!chars.includes(str[0])) {
    return str.length === 0;
  }

  return scramble(chars.replace(str[0], ''), str.slice(1));
};

const scrambleLoop = (chars, string) => {
  const normalizeStr = str => str.split('').sort().join('');

  let charsSort = normalizeStr(chars);
  const strSort = normalizeStr(string);

  for (let i = 0; i < string.length; i += 1) {
    if (!charsSort.includes(strSort[i])) {
      return false;
    }

    charsSort = charsSort.replace(strSort[i], '');
  }

  return true;
};

export { scramble, scrambleLoop };
