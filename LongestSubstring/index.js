export const longestSubstring = (s) => {
  if (s.length < 2) {
    return s.length;
  }
  let best = 0;
  let substring = [];
  for (let index = 0; index < s.length; index++) {
    const char = s[index];
    while (substring.includes(char)) {
      substring.shift();
    }
    substring.push(char);
    if (substring.length > best) {
      best = substring.length;
    }
  }
  return best;
};
