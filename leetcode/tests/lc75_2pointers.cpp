#include <iostream>
#include <format>
#include <vector>
#include <unordered_map>

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

class Solution283 {
public:
    void moveZeroes(vector<int>& nums) {
        for(size_t i = 0, j = 0; i< nums.size();)
        {
            if(nums[i] == 0)
            {
                ++i;
                continue;
            }
            if(i != j)
                swap(nums[i], nums[j]);
            ++i;
            ++j;
        }

        /* CoPiloit
        size_t j = 0;
        for (size_t i = 0; i < nums.size(); ++i)
        {
            if (nums[i] != 0)
            {
                if (i != j)
                    std::swap(nums[i], nums[j]);
                ++j;
            }
        }
         */
    }

    void test()
    {
        struct 
        {
            vector<int> nums;
            vector<int> expected;
        } test [] = {
           { {0,1,0,3,12}, {0,0,1,3,12}},
           {{0}, {0}}
        };
        cout << sizeof(test) << "\n";
        for(int i = 0; i < sizeof(test) / sizeof(test[0]) ; ++i) 
        {
            auto str = std::format("{}", to_str(test[i].nums));
            cout << str << "\n";
            moveZeroes(test[i].nums);
            cout << format("{} = {}", to_str(test[i].expected), to_str(test[i].nums)) << "\n";
        }
    }
};

class Solution392 {
public:

    bool isSubsequence(string s, string t) {
        if(s.empty())
            return true;
        if(t.empty())
            return false;
        auto i = s.begin();
        auto j = t.begin();
        while(i != s.end() && j != t.end())
        {
            if(*i == *j)
                ++i;
            ++j;
        }
        return i == s.end();
    }

    void test()
    {
        struct 
        {
            string s;
            string t;
            bool expected;
        } test [] = {
            {"abc", "ahbgdc", true},
            {"axc", "ahbgdc", false}
        };
        
        for(int i = 0; i < sizeof(test) / sizeof(test[0]) ; ++i) 
        {
            auto str = std::format("{} - {}", test[i].s, test[i].t);
            cout << str << "\n";
            auto res = isSubsequence(test[i].s, test[i].t);
            cout << format("{} = {}", test[i].expected, res) << "\n";
        }

    }
};

class Solution11{
public:
    int maxArea(vector<int>& height) {
        if(height.empty())
            return 0;
        if(height.size() == 1)
            return height[0];
        
        auto l = height.begin();
        auto r = height.end() - 1;
        int area = 0;
        while (l != r)
        {
            int a = min(*l, *r) * (r - l);
            if(area < a)
                area = a;
            if(*l < *r)
                ++l;
            else
                --r;
        }
        return area;
    }

    void test()
    {
        struct 
        {
            vector<int> height ;
            int expected;
        } test [] = {
            {{1,8,6,2,5,4,8,3,7}, 49,},
            {{1}, 1}
        };
        
        for(int i = 0; i < sizeof(test) / sizeof(test[0]) ; ++i) 
        {
            auto str = format("{} ", to_str(test[i].height));
            cout << str << "\n";
            auto res = maxArea(test[i].height);
            cout << format("{} = {}", test[i].expected, res) << "\n";
        }

    }
};

class Solution1679 {
public:
    int maxOperations(vector<int>& nums, int k) {
        unordered_map<int, int> hash;
        int count = 0;
        for(auto n : nums)
        {
            int d = k - n;
            auto& c = hash[d];
            if( c > 0)
            {
                ++count;
                c -= 1;
            }
            else
                hash[n] += 1;
            // if(auto it = hash.find(d); it != hash.end())
            // {

            // }
        }
        return count;
    }

    int maxOperations(vector<int>& nums, int k) {
        sort(nums.begin(), nums.end());
        int count = 0;
        // check sum  left and right
        return count;
    }

    void test()
    {
        struct 
        {
            vector<int> nums;
            int k;
            int expected;
        } test [] = {
            {{1,2,3,4}, 5, 2},
            {{3,1,3,4,3},6, 1}
        };
        
        for(int i = 0; i < sizeof(test) / sizeof(test[0]) ; ++i) 
        {
            cout << to_str(test[i].nums) << "\n";
            auto res = maxOperations(test[i].nums, test[i].k);
            cout << format("'{}' = '{}'", test[i].expected, res) << "\n";
        }
    }
};

int main()
{
    cout << "Begin C++ 2 pointers" << __cplusplus << "\n";
    cout << "*******************************\n";
    {
        Solution283 sol;
        sol.test();
    }
    cout << "*******************************\n";
    {
        Solution392 sol;
        sol.test();
    }
    cout << "*******************************\n";
    {
        Solution11 sol;
        sol.test();
    }
    cout << "*******************************\n";
    {
        Solution1679 sol;
        sol.test();
    }
    // cout << "*******************************\n";
    // {
    //     Solution238 sol;
    //     sol.test();
    // }
    // cout << "*******************************\n";
    // {
    //     Solution334 sol;
    //     sol.test();
    // }
    // cout << "*******************************\n";
    // {
    //     Solution53 sol;
    //     sol.test();
    // }
    // cout << "*******************************\n";
    cout << "end\n";
    return 0;
}