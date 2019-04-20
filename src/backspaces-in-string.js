// https://www.codewars.com/kata/backspaces-in-string

const processString = (str) => {
  const re = /(^|[^#])#/g;

  return str.includes('#') ? processString(str.replace(re, '')) : str;
};

export default processString;
