class Solution {
public:
    int lengthOfLongestSubstring(string s) {
        int maxLength = 0;
        
        int begin = 0;
        int curLength = 0;
        for( int i = 1; i < s.size(); i++ ) {
            int repeatIdx = repeatIndexOfChar(s, begin, i, s[i]);
            if( repeatIdx != -1 ) {
                maxLength = std::max(maxLength, i-begin);
                begin = repeatIdx + 1;
            }else{
                curLength++;
            }
        }
        int lastLength = int(s.size()) - begin;
        maxLength = std::max(maxLength, lastLength);
        return maxLength;
    }
    // Search c in range [begin, end)
    int repeatIndexOfChar(string &s, int begin, int end, char c) {
        for( int i=begin; i<end; i++ ) {
            if( s[i] == c ) {
                return i;
            }
        }
        return -1;
    }
};
