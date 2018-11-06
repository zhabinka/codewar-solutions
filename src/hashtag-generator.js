// https://www.codewars.com/kata/the-hashtag-generator

function generateHashtag(str) {
  const strProcess = str
    .split(/\s+/g)
    .filter(el => el)
    .map(el => `${el[0].toUpperCase()}${el.slice(1)}`)
    .join('');

  return (strProcess.length > 0 && strProcess.length < 140)
    ? `#${strProcess}`
    : false;
}

export default generateHashtag;
