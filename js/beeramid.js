const beeramid = (cash, price) => {
  const iter = (cans, level) => {
    const amount = cans - (level ** 2);

    return amount < 0 ? level - 1 : iter(amount, level + 1);
  };

  return iter(cash / price, 1);
};

export default beeramid;
