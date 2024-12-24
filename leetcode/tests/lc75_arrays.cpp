#include <iostream>
#include <format>
#include <vector>

using namespace std;

// https://leetcode.com/studyplan/leetcode-75/

namespace 
{
template<typename T>
    static string to_str(const vector<T>& v)
    {
        string str;
        str += '[';
        for(const auto& a : v )
        {
            str += to_string(a);
            str += ' ';
        }
        str += ']';
        return str;
    }
}

class Solution1765 {
public:
    string mergeAlternately(string word1, string word2) {
        string str;
        str.resize(word1.size() + word2.size(), 0);
        size_t i = 0, j = 0;
        for(; i < word1.size() && i < word2.size(); ++i)
        {
            str[j++] = word1[i];
            str[j++] = word2[i];
        }
        //if(word1.size() > word2.size())
            for(; i < word1.size(); ++i)
                str[j++] = word1[i];
        //else if(word1.size() < word2.size())
            for(; i < word2.size(); ++i)
                str[j++] = word2[i];
        return str;
    }

    void test()
    {
        struct 
        {
            string word1;
            string word2;
            string expected;
        } test [] = {
           { "abc", "pqr", "apbqcr"},
           { "ab", "pqrs", "apbqrs"}
        };
        cout << sizeof(test) << "\n";
        for(int i = 0; i < sizeof(test) / sizeof(test[0]) ; ++i) 
        {
            auto str = std::format("{} - {}", test[i].word1, test[i].word2);
            cout << str << "\n";
            auto res = mergeAlternately(test[i].word1, test[i].word2);
            cout << format("{} = {}", test[i].expected, res) << "\n";
        }
    }
};

class Solution1071 {
public:

    size_t gcd(size_t s1, size_t s2)
    {
        auto res = min(s1, s2);
        while(res > 0)
        {
            if(s1 % res == 0 && s2 % res == 0)
                return res;
            res -=1;
        }
        return res;
    }

    string gcdOfStrings(string str1, string str2) {
        auto cd = gcd(str1.size(), str2.size());
        for(size_t i = 0; i < cd; ++i)
        {
            if(str1[i] != str2[i])
                return {};
        }
        auto str = str1.substr(0, cd);
        for(size_t i = str.size(); i < str1.size(); ++i)
        {
            if(str1[i] != str[i % str.size()])
                return {};
        }
        for(size_t i = str.size(); i < str2.size(); ++i)
        {
            if(str2[i] != str[i % str.size()])
                return {};
        }

        return str;
    }

    void test()
    {
        struct 
        {
            string str1;
            string str2;
            string expected;
        } test [] = {
            {"ABCABC", "ABC", "ABC"},
            {"ABABAB", "ABAB", "AB"},
            {"LEET", "CODE", ""},
            {"ABCDEF", "ABC", ""}
        };
        
        for(int i = 0; i < sizeof(test) / sizeof(test[0]) ; ++i) 
        {
            auto str = std::format("{} - {}", test[i].str1, test[i].str2);
            cout << str << "\n";
            auto res = gcdOfStrings(test[i].str1, test[i].str2);
            cout << format("{} = {}", test[i].expected, res) << "\n";
        }

    }
};

class Solution1432 {
public:
    vector<bool> kidsWithCandies(vector<int>& candies, int extraCandies) {
        if(candies.empty())
            return {};
        int max = candies[0];
        for(size_t i = 1; i < candies.size(); ++i)
            if(candies[i] > max)
                max = candies[i];
        vector<bool> res;
        for(size_t i = 0; i < candies.size(); ++i)
            res.push_back(candies[i] + extraCandies >= max);
        return res;
    }

    void test()
    {
        struct 
        {
            vector<int> candies ;
            int extraCandies ;
            vector<bool> expected;
        } test [] = {
            {{2,3,5,1,3}, 3, {true,true,true,false,true},},
            {{4,2,1,1,2}, 1, {true,false,false,false,false}},
            {{12,1,12}, 10, {true,false,true}}
        };
        
        for(int i = 0; i < sizeof(test) / sizeof(test[0]) ; ++i) 
        {
            auto str = format("{} - {}", to_str(test[i].candies), test[i].extraCandies);
            cout << str << "\n";
            auto res = kidsWithCandies(test[i].candies, test[i].extraCandies);
            cout << format("{} = {}", to_str(test[i].expected), to_str(res)) << "\n";
        }

    }
};

class Solution151 {
public:
    string reverseWords(string s) {
        if(s.empty())
            return {};
        string res;
        res.reserve(s.size());
        size_t i = s.size() - 1;
        size_t end = i;
        for(; i != s.npos; --i)
        {
            if(s[i] == ' ')
            {
                if(end != i)
                {
                    if(!res.empty()) res += ' ';
                    res += s.substr(i + 1, end - i);
                }
                end = i - 1;
                continue;
            }
        }
        if(i != end)
        {
            if(!res.empty()) res += ' ';
            res += s.substr(0, end + 1);
        }
        return res;
     }

    void test()
    {
        struct 
        {
            string str1;
            string expected;
        } test [] = {
            {"EPY2giL", "EPY2giL"},
            {"the sky is blue", "blue is sky the"},
            {"  hello world  ", "world hello"},
            {"a good   example", "example good a"}
        };
        
        for(int i = 0; i < sizeof(test) / sizeof(test[0]) ; ++i) 
        {
            cout << test[i].str1 << "\n";
            auto res = reverseWords(test[i].str1);
            cout << format("'{}' = '{}'", test[i].expected, res) << "\n";
        }
    }
};

class Solution238 {
public:
    vector<int> productExceptSelf(vector<int>& nums) {
        vector<int> res;
        res.resize(nums.size());
        
        res[0] = 1;
        for(size_t i = 1; i < nums.size(); ++i)
            res[i] = res[i-1] * nums[i - 1];
        
        //cout << to_str(res) << "\n";

        int n = 1;
        for(size_t j = 0; j < nums.size(); ++j)
        {
            size_t i = nums.size() - j - 1;
            res[i] = n * res[i];
            n *= nums[i];
        }

        return res;
    }

    void test()
    {
        struct 
        {
            vector<int> nums ;
            vector<int> expected;
        } test [] = {
            {{1,2,3,4}, {24,12,8,6}},
            {{-1,1,0,-3,3}, {0,0,9,0,0}}
        };
        
        for(int i = 0; i < sizeof(test) / sizeof(test[0]) ; ++i) 
        {
            auto str = format("{}", to_str(test[i].nums));
            cout << str << "\n";
            auto res = productExceptSelf(test[i].nums);
            cout << format("{} = {}", to_str(test[i].expected), to_str(res)) << "\n";
        }

    }
};


class Solution334 {
public:
    bool increasingTriplet(vector<int>& nums) {
        int min = nums[0];
        int mid = INT_MAX;
        for(size_t i = 1; i < nums.size(); ++i)
        {
            auto n = nums[i];
            if(n <= min)
            {
                min = n;
            }
            else if(n <= mid)
            {
                mid = n;
            }
            else return true;
        }
        return false;
    }

    void test()
    {
        struct 
        {
            vector<int> nums ;
            bool expected;
        } test [] = {
            {{1,1,-2,6}, false},
            {{1,1,1,1,1,1,1,1,1,1}, false},
            {{2,4,-2,-3}, false},
            {{1,2,3,4,5}, true},
            {{5,4,3,2,1}, false},
            {{2,1,5,0,4,6}, true}
        };
        
        for(int i = 0; i < sizeof(test) / sizeof(test[0]) ; ++i) 
        {
            auto str = format("{}", to_str(test[i].nums));
            cout << str << "\n";
            auto res = increasingTriplet(test[i].nums);
            cout << format("{} = {}", test[i].expected, res) << "\n";
        }

    }
};

// https://leetcode.com/problems/maximum-subarray/description/
class Solution53 {
public:
    int maxSubArray(vector<int>& nums) {
        int sum = 0;
        int sumMax = INT_MIN;
        for(size_t i = 0; i < nums.size(); ++i)
        {
            auto n = nums[i];
            sum += n;
            sumMax = max(sum, sumMax);
            if(sum < 0)
                sum = 0;
        }
        return sumMax;
    }

    void test()
    {
        struct 
        {
            vector<int> nums ;
            int expected;
        } test [] = {
            {{-2,1,-3,4,-1,2,1,-5,4}, 6},
            {{1}, 1},
            {{5,4,-1,7,8}, 23}
        };
        
        for(int i = 0; i < sizeof(test) / sizeof(test[0]) ; ++i) 
        {
            auto str = format("{}", to_str(test[i].nums));
            cout << str << "\n";
            auto res = maxSubArray(test[i].nums);
            cout << format("{} = {}", test[i].expected, res) << "\n";
        }

    }
};

int main()
{
    cout << "Begin C++ " << __cplusplus << "\n";
    cout << "*******************************\n";
    {
        Solution1765 sol;
        sol.test();
    }
    cout << "*******************************\n";
    {
        Solution1071 sol;
        sol.test();
    }
    cout << "*******************************\n";
    {
        Solution1432 sol;
        sol.test();
    }
    cout << "*******************************\n";
    {
        Solution151 sol;
        sol.test();
    }
    cout << "*******************************\n";
    {
        Solution238 sol;
        sol.test();
    }
    cout << "*******************************\n";
    {
        Solution334 sol;
        sol.test();
    }
    cout << "*******************************\n";
    {
        Solution53 sol;
        sol.test();
    }
    cout << "*******************************\n";
    cout << "end\n";
    return 0;
}