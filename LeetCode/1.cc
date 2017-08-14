#include <vector>
#include <map>
using std::vector;
using std::map;

class Solution {
public:
    vector<int> twoSumV1(vector<int>& nums, int target) {
        typedef unsigned int uint;
        uint i = 0;
        uint j = 0;
        for(i = 0; i != (nums.size() - 1); i++) {
            int another = target - nums[i];
            for(j = i+1; j != nums.size(); j++ ) {
                if( nums[j] == another ) {
                    vector<int> result;
                    result.push_back(i);
                    result.push_back(j);
                    return result;
                }
            }
        }
        return vector<int>();
    }
    
    vector<int> twoSumV2(vector<int>& nums, int target) {
        map<int, int> numMap;
        for(unsigned int i = 0; i < nums.size(); i++) {
            int another = target - nums[i];
            map<int, int>::iterator j=numMap.find(another);
            if( j != numMap.end() ) {
                vector<int> result;
                result.push_back(j->second);
                result.push_back(i);
                return result;
            }
            numMap[nums[i]] = i;
        }
        return vector<int>();
    }
};
