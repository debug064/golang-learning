#include <iostream>
#include <format>
#include <functional>
#include <generator>
#include <map>
#include <memory>
#include <queue>
#include <sstream>
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

/**
 * Definition for a binary tree node.
 */
struct TreeNode
{
    int val;
    TreeNode *left;
    TreeNode *right;
    TreeNode() : val(0), left(nullptr), right(nullptr) {}
    TreeNode(int x) : val(x), left(nullptr), right(nullptr) {}
    TreeNode(int x, TreeNode *left, TreeNode *right) : val(x), left(left), right(right) {}
};

std::generator<string> split(string_view str)
{
    size_t b = 0;
    queue<TreeNode*> q1, q2;
    while (b < str.size())
    {
        auto e = str.find(',', b);
        if(e == str.npos)
            e = str.size();
        auto s = str.substr(b, e - b);
        b = e + 1;
        auto res = s == "null" ? "" : string(s);
        co_yield res;
    }
}

unique_ptr<TreeNode> stringToTree(string_view str)
{
    auto vals = split(str);
    auto it = vals.begin();
    if (it == vals.end())
        return {};
    unique_ptr<TreeNode> res;
    res = make_unique<TreeNode>(stoi(*it));
    ++it;
    queue<TreeNode*> q1, q2;
    q1.push(res.get());
    while(!q1.empty())
    {
        auto *node = q1.front();
        q1.pop();
        if(it == vals.end())
            break;
        if(*it != "")
        {
            node->left = new TreeNode(stoi(*it));
            q2.push(node->left);
        }
        ++it;
        if(it == vals.end())
            break;
        if(*it != "")
        {
            node->right = new TreeNode(stoi(*it));
            q2.push(node->right);
        }
        if(q1.empty())
            swap(q1, q2);
        ++it;
    }
    return res;
}

string treeToStr(TreeNode* root) 
{
    if(root == nullptr)
        return {};
    stringstream str;
    queue<TreeNode*> q1, q2;
    str << to_string(root->val);
    q1.push(root->left);
    q1.push(root->right);
    while(!q1.empty())
    {
        if(q1.front())
        {
            str << "," << to_string(q1.front()->val);
            q2.push(q1.front()->left);
            q2.push(q1.front()->right);
        }
        else
        {
            str << ",null";
        }
        q1.pop();
        if(q1.empty())
            swap(q1, q2);
    }
    auto res = str.str();
    auto e = res.size();
    while(e >=5 && res.compare(e - 5, 5, ",null") == 0)
        e -= 5;
    return res.substr(0, e);
}

class Solution700 {
public:
    TreeNode *searchBST(TreeNode *root, int val)
    {
        if (root == nullptr)
            return nullptr;
        if (root->val == val)
            return root;
        if (val < root->val)
            return searchBST(root->left, val);
        else
            return searchBST(root->right, val);
    }

    void test()
    {
        struct 
        {
            string root1;
            int val;
            string expected;
        } test [] = {
            {"4,2,7,1,3", 2, "2,1,3"},
            {"4,2,7,1,3", 5 ,""},
            //{ {"3,2,5,null,null,4,10,null,null,8,15,7"}, {"3,2,7,null,null,4,10,null,null,8,15"}},
        };
        cout << sizeof(test) << "\n";
        for(int i = 0; i < sizeof(test) / sizeof(test[0]) ; ++i) 
        {
            cout << test[i].root1 << " : " << test[i].val << "\n";
            auto root = stringToTree(test[i].root1);
            auto res = searchBST(root.get(), test[i].val);
            auto str = treeToStr(res);
            cout << format("{} = {}", test[i].expected, str) << "\n";
        }
    }
};

class Solution450 {
public:
    TreeNode* deleteNode(TreeNode* root, int key) {
        if (root == nullptr)
            return nullptr;
        if (key < root->val)
            root->left = deleteNode(root->left, key);
        else if (key > root->val)
            root->right = deleteNode(root->right, key);
        else
        {
            if(root->left == nullptr)
            {
                auto res = root->right;
                delete root;
                return res;
            }
            else if(root->right == nullptr)
            {
                auto res = root->left;
                delete root;
                return res;
            }
            else
            {
                auto r = root->right;
                while(r->left)
                    r = r->left;
                root->val = r->val;
                root->right = deleteNode(root->right, r->val);
            }
            
        }
        return root;
    }

    void test()
    {
        struct 
        {
            string root1;
            int val;
            string expected;
        } test [] = {
            {"3,2,5,null,null,4,10,null,null,8,15,7", 5, "3,2,7,null,null,4,10,null,null,8,15"},
            {"5,3,6,2,4,null,7", 3, "5,4,6,2,null,null,7"},
            {"50,30,70,null,40,60,80", 50 ,"60,30,70,null,40,null,80"},
            {"0", 0 ,""},
            {"5,3,6,2,4,null,7", 7, "5,3,6,2,4"},
            {"5,3,6,2,4,null,7", 0 ,"5,3,6,2,4,null,7"},
            
        };

        for(int i = 0; i < sizeof(test) / sizeof(test[0]) ; ++i) 
        {
            cout << format("{} - {}", test[i].root1, test[i].val) << "\n";
            auto root = stringToTree(test[i].root1);
            auto res = deleteNode(root.get(), test[i].val);
            if(res == nullptr)
                root.release();
            auto str = treeToStr(res);
            cout << format("{} = {}", test[i].expected, str) << "\n";
        }
    }
};

class Solution450_2 {
    public:
        TreeNode* deleteNode(TreeNode* root, int key) {
            if (root == nullptr)
                return nullptr;
            if(key < root->val)
                root->left = deleteNode(root->left, key);
            else if (key > root->val)
                root->right = deleteNode(root->right, key);
            else {
                if (root->left == nullptr) {
                    auto tmp = root->right;
                    delete root;
                    root = tmp;
                } else if (root->right == nullptr) {
                    auto tmp = root->left;
                    delete root;
                    root = tmp;
                }
                else {
                    auto tmp = root->right;
                    while (tmp->left)
                        tmp = tmp->left;
                    root->val = tmp->val;
                    root->right = deleteNode(root->right, root->val);
                }
            }
            
            return root;
        }
    
        void test()
        {
            struct 
            {
                string root1;
                int val;
                string expected;
            } test [] = {
                {"3,2,5,null,null,4,10,null,null,8,15,7", 5, "3,2,7,null,null,4,10,null,null,8,15"},
                {"5,3,6,2,4,null,7", 3, "5,4,6,2,null,null,7"},
                {"50,30,70,null,40,60,80", 50 ,"60,30,70,null,40,null,80"},
                {"0", 0 ,""},
                {"5,3,6,2,4,null,7", 7, "5,3,6,2,4"},
                {"5,3,6,2,4,null,7", 0 ,"5,3,6,2,4,null,7"},
                
            };
    
            for(int i = 0; i < sizeof(test) / sizeof(test[0]) ; ++i) 
            {
                cout << format("{} - {}", test[i].root1, test[i].val) << "\n";
                auto root = stringToTree(test[i].root1);
                auto res = deleteNode(root.get(), test[i].val);
                if(res == nullptr)
                    root.release();
                auto str = treeToStr(res);
                cout << format("{} = {}", test[i].expected, str) << "\n";
            }
        }
    };
    


int main()
{
    cout << "Begin C++ " << __cplusplus << "\n";
    cout << "*******************************\n";
    {
        Solution700 sol;
        sol.test();
    }
    cout << "\n*******************************\n";
    {
        Solution450 sol;
        sol.test();
    }
    //cout << format("{:*>{}}\n", "", 25) ;
    cout << format("\n{:*^35}\n", " Solution450_2 ") ;
    {
        Solution450_2 sol;
        sol.test();
    }
    cout << format("{:*^45}\n", "") ;
    cout << "end\n";
    return 0;
}