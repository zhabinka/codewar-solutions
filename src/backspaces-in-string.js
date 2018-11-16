const processString = (str) => {
  const re = /\w?#/;

  return re.test(str) ? processString(str.replace(re, '')) : str;
};

export default processString;
