--lua内部全局变量（保留字）
print(_VERSION)

--nil
a = 1
print("a:", a)
a = nil
print("delete a:", a)

ntbl = { k1 = "v1", k2 = "v2" }
print("table.k1:", ntbl.k1)
ntbl.k1 = nil
print("delete table.k1:", ntbl.k1)
print("table.k2:", ntbl.k2)

print(type(x) == "nil") --true
print(type(x) == nil) --false

--boolean
if false or nil then
    print("at least one true")
else
    print("false and nil are both false")
end

if 0 then
    print("number 0 is true")
else
    print("number 0 is false")
end

--string，索引从1开始
str1 = "this is a str1"
str2 = 'this s str2'
str3 = [[this is str3]]
print(str1)
print(str2)
print(str3)
print("1" + 2) --3
print('2e2' * 3) --600
print("a" .. "b") --ab
print(123 .. 456) --123456

--table，就是关联表，key可以为数字(负数、0、正数都行)或字符串，value是什么都行。默认索引从1开始，table会自动扩容
tb2 = {} --nil
tb1 = { 1, "a", k1 = "v1" } ---用这种方式给key赋值时，key不能用引号包起来
tb1.k2 = "k2" --用这种方式给key赋值时，key不能用引号包起来
tb1["k3"] = "v3"    --用这种方式给key赋值时，key必须用引号包起来
k4 = "k4"
tb1[k4] = "v4"  --等价于tb1["k4"]="v4"
tb1[100] = 666
for k, v in pairs(tb1) do
    print("key:" .. k .. ",val:" .. v)
end

--function，在 Lua 中，函数是被看作是"第一类值（First-Class Value）"，可以存在变量里
local function factorial1(n)
    if n == 0 then
        return 1
    else
        return n * factorial1(n - 1)
    end
end
print(factorial1(5))
local factorial2 = factorial1
print(factorial2(5))

--将函数以匿名函数的方式传参
function anonymous(tab, fun)
    for k, v in pairs(tab) do
        print(fun(k, v))
    end
end
tab = { key1 = "val1", key2 = "val2" }
anonymous(tab, function(key, val)
    return key .. " = " .. val
end)

--变量交换，和Go一样，先把右边的计算好，然后分别赋值给坐标
x, y = y, x
--多个变量赋值，值不够的为nil，值多了的忽略
a1, b1, c1 = 1, 2 --c1为nil
a2, b2 = 1, 2, 3 --3被忽略

--while循环
local a = 1
while (a < 5) do
    a = a + 1
    print("while loop, a:" .. a)
end

--数值for循环
for i = 2, 10, 3 do
    print("for value loop" .. i)
end

--泛型for循环
local aa = { "cao" }
aa[3] = "bo"
aa[5] = "ri"
aa[-1] = "en"

for k, v in pairs(aa) do
    --pairs能遍历tale中的所有元素，即使索引<=0
    print("pairs" .. k .. v)
end

for k, v in ipairs(aa) do
    --ipairs碰到nil就停止，一般用来遍历数组
    print("ipairs" .. k .. v)
end

--repeat until循环
local rep = 5
repeat
    print("repeat until:" .. a)
    a = a + 1
until (a > 10)

--if else if else很简单，没啥可说的

--函数
function max(num1, num2)
    if (num1 > num2) then
        result = num1;
    else
        result = num2;
    end
    return result;
end
print("max func:" .. max(10, -2))
--局部函数
local function exchange(num1, num2)
    if (num1 < num2) then
        return num1, num2;
    else
        return num2, num1;
    end
end
print("exchange func:" .. exchange(5, 9))
--可变参数函数，可以把可变参数放入table中
local function average(...)
    result = 0
    local arr = { ... }
    for _, v in ipairs(arr) do
        result = result + v
    end
    print("total " .. #arr .. " params")
    return result / #arr
end
print("average func:" .. average(1, 2, 5, 8, 10))

--[[
 运算符：+ - * / % ^ - == ~= > < >= <= and or not .. #
 和Go不太一样的：三个逻辑运算符；~=是不等于；..是连接字符串；#是返回字符串或表的长度
]]--

--字符串的三种定义方式
local str1 = 'this is str1'
local str2 = "this is str2"
local str3 = [[this is str3]]

--[[
 转义字符：\a \b \f \n \r \t \v \\ \' \" ，和Go一样
 \ddd 是1到3位八进制数所代表的任意字符
 \xhh 是1到2位十六进制所代表的任意字符
]]--

--字符串函数
--upper：字符串全部转为大写字母
print("upper:" .. string.upper("AbC"))
--lower：字符串全部转为小写字母
print("lower:" .. string.lower("AbC"))
--[[
 gsub：在字符串中替换
 第一个参数为要替换的字符串
 第二个参数为被替换的字符
 第三个参数为替换成什么字符
 第四个参数为替换次数（可以忽略，则全部替换）
]]--
print("gsub:" .. string.gsub("ababab", "a", "s", 2))
--[[
 find：查找指定字符串
 第一个参数为原字符串
 第二个参数为要匹配的字符串
 第三个参数为索引(正数表示从第几位开始匹配，负数则是从倒数第几位开始匹配)
 第四个参数默认为false，为true函数只默认简单查找子串(即把%d %f %s %c认为是普通的字符串)，为false函数会按模式匹配查找(即把%d %f %s %c等转换成占位符)
]]
print("find:" .. string.find("a3b%dcdfg%d", "%d", 4, true))
print("find:" .. string.find("a3b%dcdfg%d", "%d", 1, false))
--字符串反转
print("reverse:" .. string.reverse("abc"))
--格式化字符串，和其他语言的printf一样
print("format:" .. string.format("test:%05d", 1))
--将整型数字转成字符并连接
print("char:" .. string.char(99, 33, 97))
--转换字符为整数值，第二个参数为指定的字符，省略则默认第一个字符
print("byte:" .. string.byte('abcdef', 3))
--计算字符串长度，等价于#
print("len:" .. string.len("abc"))
print("#:" .. #"abc")
--返回字符串的n个拷贝
print("rep:" .. string.rep("ab", 5))

--无状态迭代器实现ipairs
function iter (it, i)
    --迭代器
    i = i + 1
    if it[i] then
        --非nil和false即为true
        return i, it[i]
    end
end

function myipairs(it)
    --返回迭代器、状态常量、控制常量的函数
    return iter, it, 0
end
aa = {}
for i = -5, 5 do
    --给数组a赋值
    aa[i] = i * 2
end

for i, v in pairs(aa) do
    print("pairs", i, v)
end

for i, v in myipairs(aa) do
    --使用无状态迭代器方式1
    print("mypairs", i, v)
end

for i, v in iter, aa, 1 do
    --使用无状态迭代器方式2
    print("iter", i, v)
end

--多状态迭代器，把迭代器的信息封装到table内，且用闭包函数实现迭代函数
array = { "Lua", "Tutorial" }
function elementIterator (collection)
    local index = 0
    local count = #collection
    -- 闭包函数
    return function()
        index = index + 1
        if index <= count
        then
            --  返回迭代器的当前元素
            return collection[index]
        end
    end
end

for element in elementIterator(array)
do
    print(element)
end

--table.concat
t = { "wo", "ri", "ni", "ma" }
print(table.concat(t))
print(table.concat(t, ":"))
print(table.concat(t, "#", 2))
print(table.concat(t, "@", 2, 3))

--table.insert
table.insert(t, "?")
print("after insert:" .. t[#t])
table.insert(t, 2, "!")
print("after insert:" .. t[2])

--table.remove
print("before remove:" .. t[#t])
table.remove(t)
print("after remove:" .. t[#t])

print("before remove:" .. t[2])
table.remove(t, 2)
print("after remove:" .. t[2])

--table.sort
table.sort(t)
for i, v in ipairs(t) do
    print("sorted:" .. v)
end

--module
require("mymodule")   --或者 require<"mymodule">
print(mymodule.const1)
print(mymodule["const2"])
mymodule.function1()
mymodule.function3()

local m = require("mymodule") --给加载的模块定义别名，然后可通过别名调用模块的成员
m.function3()
