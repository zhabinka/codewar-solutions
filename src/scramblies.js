// https://www.codewars.com/kata/scramblies/

const scramble = (chars, str) => {
  return [...str].reduce((acc, el) => (acc ? chars.includes(el) : false));
};

export default scramble;
