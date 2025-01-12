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

class Solution643 {
public:
    double findMaxAverage(vector<int>& nums, int k) {
        int res = INT_MIN;
        int m = 0;
        int j = 0;
        for(size_t i = 0; i < nums.size(); ++i)
        {
            m += nums[i];
            ++j;
            if(j == k)
            {
                if(m > res)
                    res = m;
                m -= nums[i - k + 1];
                --j;
            }
        }
        return (double)res / k;
    }

    void test()
    {
        struct 
        {
            vector<int> nums ;
            int k ;
            double expected;
        } test [] = {
           { {1,12,-5,-6,50,3}, 4, 12.75},
           { {5}, 1, 5.}
        };
        cout << sizeof(test) << "\n";
        for(int i = 0; i < sizeof(test) / sizeof(test[0]) ; ++i) 
        {
            auto str = std::format("{} - {}", to_str(test[i].nums), test[i].k);
            cout << str << "\n";
            auto res = findMaxAverage(test[i].nums, test[i].k);
            cout << format("{} = {}", test[i].expected, res) << "\n";
        }
    }
};

class Solution1456 {
public:


    int maxVowels(string s, int k) {
        auto isVowel = [](char c) {
            return c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u';
        };
        int res = 0;
        int m = 0;
        int j = 0;
        for(size_t i = 0; i < s.size(); ++i)
        {
            if(isVowel(s[i]))
                ++m;
            ++j;
            if( m > res)
                res = m;
            if(j == k)
            {
                --j;
                if(isVowel(s[i - j ]))
                    --m;
            }
        }
        return res;
    }

    void test()
    {
        struct 
        {
            string s;
            int k;
            int expected;
        } test [] = {
            {"abciiidef", 3, 3},
            {"aeiou", 2, 2},
            {"leetcode", 3, 2}
        };
        
        for(int i = 0; i < sizeof(test) / sizeof(test[0]) ; ++i) 
        {
            auto str = std::format("{} - {}", test[i].s, test[i].k);
            cout << str << "\n";
            auto res = maxVowels(test[i].s, test[i].k);
            cout << format("{} = {}", test[i].expected, res) << "\n";
        }

    }
};

class Solution1004 {
public:
    int longestOnes(vector<int>& nums, int k) {
        size_t i(0), j (0);
        int zeros = 0;
        int res = 0;
        int m = 0;
        for( ; i < nums.size(); ++i)
        {
            if(nums[i] == 1)
                ++m;
            else
            {
                ++zeros;
                ++m;
                if(zeros > k)
                {
                    for( ; j <= i; ++j)
                    {
                        if(m)
                        --m;
                        if(nums[j] == 0)
                        {
                            --zeros;
                            ++j;
                            break;
                        }
                    }
                }
            }
            if(res < m)
                res = m;
        }
        return res;
    }

    void test()
    {
        struct 
        {
            vector<int> nums ;
            int k ;
            int expected;
        } test [] = {
            {{1,1,1,0,0,0,1,1,1,1}, 0, 4},
            {{0,0,1,1,1,0,0}, 0, 3},
            {{0,1,1,1,0,0,0,1,1,1,1,0}, 2, 6},
            {{0,0,1,1,0,0,1,1,1,0,1,1,0,0,0,1,1,1,1}, 3, 10}
        };
        
        for(int i = 0; i < sizeof(test) / sizeof(test[0]) ; ++i) 
        {
            auto str = format("{} - {}", to_str(test[i].nums), test[i].k);
            cout << str << "\n";
            auto res = longestOnes(test[i].nums, test[i].k);
            cout << format("{} = {}", test[i].expected, res) << "\n";
        }

    }
};

class Solution1493 {
public:
    int longestSubarray(vector<int>& nums) {
        int res = 0, m = 0;
        size_t z = 0, r = 0, l = 0;
        for(; r < nums.size(); ++r)
        {
            if(nums[r] == 0)
            {
                l = z;
                if(nums[z] == 0)
                    ++l;
                z = r;
                
            }
            m = r -z;
            if (z > l)
                m += z - l;
            if(m > res)
                res = m;
        }
        return res;
    }

    int longestSubarray2(vector<int>& nums) {
        size_t z = 0, l = 0, r;
        int m = 0;
        int res = 0;
        for(size_t r = 0; r < nums.size(); ++r) {
            if(nums[r] == 1) {
                ++m;
                if(res < m)
                    res = m;
                continue;
            }
            for(; l < z; ++l) {
                if(nums[l] == 1)
                    --m;
            }
            z = r;
        }
        return res - nums[z];   
    }

    void test()
    {
        struct 
        {
            vector<int> nums ;
            int expected;
        } test [] = {
            {{0,1,1,1,1,1}, 5},
            {{0,1,1,1,0,1,1,0,1}, 5},
            {{1,1,0,1}, 3},
            {{0,0,0,1,1,1,1,1}, 5},
            {{1,1,1}, 2}
        };
        
        for(int i = 0; i < sizeof(test) / sizeof(test[0]) ; ++i) 
        {
            cout << to_str(test[i].nums) << "\n";
            auto res = longestSubarray(test[i].nums);
            cout << format("'{}' = '{}'", test[i].expected, res) << "\n";
        }
        cout << "------------------------\n";
        for(int i = 0; i < sizeof(test) / sizeof(test[0]) ; ++i) 
        {
            cout << to_str(test[i].nums) << "\n";
            auto res = longestSubarray2(test[i].nums);
            cout << format("'{}' = '{}'", test[i].expected, res) << "\n";
        }
    }
};


int main()
{
    cout << "Begin C++ " << __cplusplus << "\n";
    cout << "*******************************\n";
    {
        Solution643 sol;
        sol.test();
    }
    cout << "*******************************\n";
    {
        Solution1456 sol;
        sol.test();
    }
    cout << "*******************************\n";
    {
        Solution1004 sol;
        sol.test();
    }
    cout << "*******************************\n";
    {
        Solution1493 sol;
        sol.test();
    }
    cout << "*******************************\n";

    cout << "end\n";
    return 0;
}