import { longestSubstring } from '.';

describe('Longest Substring', () => {
  it('returns the longest substring', () => {
    const testCases = [
      {
        input: 'abcabcbb',
        output: 3,
      },
      {
        input: 'bbbbb',
        output: 1,
      },
      {
        input: 'pwwkew',
        output: 3,
      },
      {
        input: '',
        output: 0,
      },
      {
        input: 'abba',
        output: 2,
      },
      {
        input: 'dvdf',
        output: 3,
      },
    ];
    testCases.forEach(({ input, output }) => {
      const received = longestSubstring(input);
      expect(received).toBe(output);
    });
  });
});
