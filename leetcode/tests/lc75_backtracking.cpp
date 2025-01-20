#include <iostream>
#include <format>
#include <functional>
#include <map>
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

    template<>
     string to_str(const vector<string>& v)
    {
        string str;
        str += '[';
        for(const auto& a : v )
        {
            str += a;
            str += ' ';
        }
        str += ']';
        return str;
    }

    template<typename T>
    static string to_str(const vector<vector<T>>& v)
    {
        string str;
        str += '[';
        for(const auto& a : v )
        {
            str += to_str(a);
            str += ' ';
        }
        str += ']';
        return str;
    }
}

class Solution17 {
public:
    vector<string> letterCombinations(string digits)
    {
        map<int, string> phone = {
            {2, "abc"},
            {3, "def"},
            {4, "ghi"},
            {5, "jkl"},
            {6, "mno"},
            {7, "pqrs"},
            {8, "tuv"},
            {9, "wxyz"}};

        vector<string> res;
        std::function<void(string)> backtr = [&](string nums)
        {
            if (nums.empty())
                return;
            if(nums.size() == 1)
            {
                auto str = phone[atoi(nums.c_str())];
                for(auto c : str)
                    res.push_back(string(1, c));
                return;
            }
            backtr(nums.substr(1));
            vector<string> tmp;
            auto str = phone[atoi(nums.substr(0,1).c_str())];
            for(auto c : str)
                for(auto s : res)
                    tmp.push_back(string(1, c) + s);
            res = tmp;
        };

        backtr(digits);

        return res;
    }

    void test()
    {
        struct 
        {
            string digits;
            vector<string> expected;
        } test [] = {
           { "23", {"ad","ae","af","bd","be","bf","cd","ce","cf"}},
           { "2", {"a","b","c"}},
           { "", {}},
        };
        cout << sizeof(test) << "\n";
        for(int i = 0; i < sizeof(test) / sizeof(test[0]) ; ++i) 
        {
            auto str = std::format("{}", test[i].digits);
            cout << str << "\n";
            auto res = letterCombinations(test[i].digits);
            cout << format("{} = {}", to_str(test[i].expected), to_str(res)) << "\n";
        }
    }
};

class Solution216 {
public:
    vector<vector<int>> combinationSum3(int k, int sum) 
    {
        vector<vector<int>> res;
        
        function<void(int, int,int, vector<int>)> back_sum = [&](int beg, int k, int sum, vector<int> pass)
        {
            if( k < 0)
                return;
            if(sum < 0)
                return;
            if(k == 0 && sum == 0)
            {
                res.push_back(pass);
                return;
            }
            for(int i = beg; i <= 9; ++i )
            {
                pass.push_back(i);
                back_sum(i + 1, k-1, sum - i, pass);
                pass.pop_back();
            }
        };
        back_sum(1, k, sum, {});
        
        return res;
    }

    void test()
    {
        struct 
        {
            int k ;
            int n;
            vector<vector<int>> expected;
        } test [] = {
            {3, 7, {{1,2,4}}},
            {3, 9, {{1, 2, 6}, {1, 3, 5}, {2, 3, 4}}},
            {4, 1, {}}
        };
        
        for(int i = 0; i < sizeof(test) / sizeof(test[0]) ; ++i) 
        {
            auto str = format("{} - {}", test[i].k, test[i].n);
            cout << str << "\n";
            auto res = combinationSum3(test[i].k, test[i].n);
            cout << format("{} = {}", to_str(test[i].expected), to_str(res)) << "\n";
        }

    }
};



int main()
{
    cout << "Begin C++ " << __cplusplus << "\n";
    cout << "*******************************\n";
    {
        Solution17 sol;
        sol.test();
    }
    cout << "*******************************\n";
    {
        Solution216 sol;
        sol.test();
    }
    cout << "*******************************\n";
    cout << "end\n";
    return 0;
}