const filterAnagrams = (word, words) => {
  const normalizeStr = str => [...str].sort().join('');

  return words.filter(el => normalizeStr(el) === normalizeStr(word));
};

export default filterAnagrams;
