/*
 * 程序名：book205.cpp，此程序演示用类的基本概念。
 * 作者：C语言技术网(www.freecplus.net) 日期：20190525
*/
#include <iostream>
using namespace std;
#include <stdio.h>
#include <string.h>
 
class CGirl    // 定义超女类
{
public:
  char m_name[50];  // 姓名
  int  m_age;       // 年龄
  int  m_height;    // 身高，单位：厘米cm
  char m_sc[30];    // 身材，火辣；普通；飞机场。
  char m_yz[30];    // 颜值，漂亮；一般；歪瓜裂枣。
  int  Show();      // 申明显示超女基本信息的成员函数
};
 
int main()
{
  CGirl Girl;   // 实例化一个Girl对象
 
  // 访问对象的成员变量，进行赋值
  strcpy(Girl.m_name,"武则天");
  Girl.m_age=28;     
  Girl.m_height=168;
  strcpy(Girl.m_sc,"火辣");
  strcpy(Girl.m_yz,"漂亮");
 
  Girl.Show();   // 调用对象的成员函数
  cout <<"1111"<<endl;
}
 
int CGirl::Show()  // 显示超女基本信息的成员函数体
{
  printf("姓名：%s，年龄：%d，身高：%d，身材：%s，颜值：%s。\n",m_name,m_age,m_height,m_sc,m_yz);
}
