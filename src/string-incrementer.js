const separate = str => [str.match(/^\D+/), str.match(/\d*/g).join('')];

const strInc = (str) => {
  const [chars, num] = separate(str);
  const inc = (+num + 1)
    .toString()
    .padStart(num.length, 0);

  return chars
    .concat(inc)
    .join('');
};

export default strInc;
