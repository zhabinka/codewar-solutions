// https://www.codewars.com/kata/double-cola/

const whoIsNext = (names, cans) => {
  const iter = (list, n, eque) => {
    for (let i = 0; i < list.length; i += 1) {
      if (n - eque * (i + 1) <= 0) {
        return list[i];
      }
    }

    return iter(list, n - list.length * eque, eque * 2);
  };

  return iter(names, cans, 1);
};

export default whoIsNext;
