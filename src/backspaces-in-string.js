const processString = (str) => {
  const re = /.?#/;

  return re.test(str) ? processString(str.replace(re, '')) : str;
};

export default processString;
