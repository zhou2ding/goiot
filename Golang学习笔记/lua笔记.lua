--lua�ڲ�ȫ�ֱ����������֣�
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

--string��������1��ʼ
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

--table�����ǹ�����key����Ϊ����(������0����������)���ַ�����value��ʲô���С�Ĭ��������1��ʼ��table���Զ�����
tb2 = {} --nil
tb1 = { 1, "a", k1 = "v1" } ---�����ַ�ʽ��key��ֵʱ��key���������Ű�����
tb1.k2 = "k2" --�����ַ�ʽ��key��ֵʱ��key���������Ű�����
tb1["k3"] = "v3"    --�����ַ�ʽ��key��ֵʱ��key���������Ű�����
k4 = "k4"
tb1[k4] = "v4"  --�ȼ���tb1["k4"]="v4"
tb1[100] = 666
for k, v in pairs(tb1) do
    print("key:" .. k .. ",val:" .. v)
end

--function���� Lua �У������Ǳ�������"��һ��ֵ��First-Class Value��"�����Դ��ڱ�����
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

--�����������������ķ�ʽ����
function anonymous(tab, fun)
    for k, v in pairs(tab) do
        print(fun(k, v))
    end
end
tab = { key1 = "val1", key2 = "val2" }
anonymous(tab, function(key, val)
    return key .. " = " .. val
end)

--������������Goһ�����Ȱ��ұߵļ���ã�Ȼ��ֱ�ֵ������
x, y = y, x
--���������ֵ��ֵ������Ϊnil��ֵ���˵ĺ���
a1, b1, c1 = 1, 2 --c1Ϊnil
a2, b2 = 1, 2, 3 --3������

--whileѭ��
local a = 1
while (a < 5) do
    a = a + 1
    print("while loop, a:" .. a)
end

--��ֵforѭ��
for i = 2, 10, 3 do
    print("for value loop" .. i)
end

--����forѭ��
local aa = { "cao" }
aa[3] = "bo"
aa[5] = "ri"
aa[-1] = "en"

for k, v in pairs(aa) do
    --pairs�ܱ���tale�е�����Ԫ�أ���ʹ����<=0
    print("pairs" .. k .. v)
end

for k, v in ipairs(aa) do
    --ipairs����nil��ֹͣ��һ��������������
    print("ipairs" .. k .. v)
end

--repeat untilѭ��
local rep = 5
repeat
    print("repeat until:" .. a)
    a = a + 1
until (a > 10)

--if else if else�ܼ򵥣�ûɶ��˵��

--����
function max(num1, num2)
    if (num1 > num2) then
        result = num1;
    else
        result = num2;
    end
    return result;
end
print("max func:" .. max(10, -2))
--�ֲ�����
local function exchange(num1, num2)
    if (num1 < num2) then
        return num1, num2;
    else
        return num2, num1;
    end
end
print("exchange func:" .. exchange(5, 9))
--�ɱ�������������԰ѿɱ��������table��
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
 �������+ - * / % ^ - == ~= > < >= <= and or not .. #
 ��Go��̫һ���ģ������߼��������~=�ǲ����ڣ�..�������ַ�����#�Ƿ����ַ������ĳ���
]]--

--�ַ��������ֶ��巽ʽ
local str1 = 'this is str1'
local str2 = "this is str2"
local str3 = [[this is str3]]

--[[
 ת���ַ���\a \b \f \n \r \t \v \\ \' \" ����Goһ��
 \ddd ��1��3λ�˽�����������������ַ�
 \xhh ��1��2λʮ������������������ַ�
]]--

--�ַ�������
--upper���ַ���ȫ��תΪ��д��ĸ
print("upper:" .. string.upper("AbC"))
--lower���ַ���ȫ��תΪСд��ĸ
print("lower:" .. string.lower("AbC"))
--[[
 gsub�����ַ������滻
 ��һ������ΪҪ�滻���ַ���
 �ڶ�������Ϊ���滻���ַ�
 ����������Ϊ�滻��ʲô�ַ�
 ���ĸ�����Ϊ�滻���������Ժ��ԣ���ȫ���滻��
]]--
print("gsub:" .. string.gsub("ababab", "a", "s", 2))
--[[
 find������ָ���ַ���
 ��һ������Ϊԭ�ַ���
 �ڶ�������ΪҪƥ����ַ���
 ����������Ϊ����(������ʾ�ӵڼ�λ��ʼƥ�䣬�������Ǵӵ����ڼ�λ��ʼƥ��)
 ���ĸ�����Ĭ��Ϊfalse��Ϊtrue����ֻĬ�ϼ򵥲����Ӵ�(����%d %f %s %c��Ϊ����ͨ���ַ���)��Ϊfalse�����ᰴģʽƥ�����(����%d %f %s %c��ת����ռλ��)
]]
print("find:" .. string.find("a3b%dcdfg%d", "%d", 4, true))
print("find:" .. string.find("a3b%dcdfg%d", "%d", 1, false))
--�ַ�����ת
print("reverse:" .. string.reverse("abc"))
--��ʽ���ַ��������������Ե�printfһ��
print("format:" .. string.format("test:%05d", 1))
--����������ת���ַ�������
print("char:" .. string.char(99, 33, 97))
--ת���ַ�Ϊ����ֵ���ڶ�������Ϊָ�����ַ���ʡ����Ĭ�ϵ�һ���ַ�
print("byte:" .. string.byte('abcdef', 3))
--�����ַ������ȣ��ȼ���#
print("len:" .. string.len("abc"))
print("#:" .. #"abc")
--�����ַ�����n������
print("rep:" .. string.rep("ab", 5))

--��״̬������ʵ��ipairs
function iter (it, i)
    --������
    i = i + 1
    if it[i] then
        --��nil��false��Ϊtrue
        return i, it[i]
    end
end

function myipairs(it)
    --���ص�������״̬���������Ƴ����ĺ���
    return iter, it, 0
end
aa = {}
for i = -5, 5 do
    --������a��ֵ
    aa[i] = i * 2
end

for i, v in pairs(aa) do
    print("pairs", i, v)
end

for i, v in myipairs(aa) do
    --ʹ����״̬��������ʽ1
    print("mypairs", i, v)
end

for i, v in iter, aa, 1 do
    --ʹ����״̬��������ʽ2
    print("iter", i, v)
end

--��״̬���������ѵ���������Ϣ��װ��table�ڣ����ñհ�����ʵ�ֵ�������
array = { "Lua", "Tutorial" }
function elementIterator (collection)
    local index = 0
    local count = #collection
    -- �հ�����
    return function()
        index = index + 1
        if index <= count
        then
            --  ���ص������ĵ�ǰԪ��
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
require("mymodule")   --���� require<"mymodule">
print(mymodule.const1)
print(mymodule["const2"])
mymodule.function1()
mymodule.function3()

local m = require("mymodule") --�����ص�ģ�鶨�������Ȼ���ͨ����������ģ��ĳ�Ա
m.function3()
