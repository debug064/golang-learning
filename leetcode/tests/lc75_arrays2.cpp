#include <climits>
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


    string gcdOfStrings(string str1, string str2) {
        string str;
    
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
        vector<bool> res;
        
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
        string res;
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