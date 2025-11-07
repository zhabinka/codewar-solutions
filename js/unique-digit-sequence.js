// https://www.codewars.com/kata/unique-digit-sequence/

Array.prototype.last = function last() {
  return this[this.length - 1];
};

const getNextFree = (arr, i) => (arr.includes(i) ? getNextFree(arr, i + 1) : i);

const getNextValue = (sequence, i = 0) => {
  const value = getNextFree(sequence, i);
  const lastValueDigit = [...sequence.last().toString()];
  const valueDigit = [...value.toString()];
  const intersectionDigit = lastValueDigit.filter(el => valueDigit.includes(el));

  return intersectionDigit.length > 0
    ? getNextValue(sequence, value + 1)
    : value;
};

const sequenceBuild = (n) => {
  const iter = (arr, i) => {
    if (i > n) {
      return arr;
    }

    return iter(arr.concat(getNextValue(arr)), i + 1);
  };

  return iter([0], 1);
};

const findNum = index => sequenceBuild(index)[index];

export default findNum;
