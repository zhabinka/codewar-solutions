const processString = (str) => {
  const re = /(^|[^#])#/g;

  return str.includes('#') ? processString(str.replace(re, '')) : str;
};

export default processString;
