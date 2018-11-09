// https://www.codewars.com/kata/scramblies/

const scramble = (chars, str) => {
  if (!chars.includes(str[0])) {
    return str.length === 0;
  }

  return scramble(chars.replace(str[0], ''), str.slice(1));
};

export default scramble;
