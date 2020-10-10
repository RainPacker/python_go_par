#include <stdio.h>
#include <math.h>
#include <string.h>
#define URL "http://www.fishc.com"
#define NAME "杨杨"


void print_mulit_table(){
    int i=1;
    while (i<=9)
    {   
        int j=1;
        while (j<=i)
        {
            printf("%d*%d=%d\t",j, i, i*j);
            j++;
        }
        
        i++;
        printf("\n");
    }
    
}
void str_cat(){
    char src[50], dest[50];
 
   strcpy(src,  "This is source");
   strcpy(dest, "This is destination");
 
   strcat(dest, src);
 
   printf("最终的目标字符串： |%s|", dest);
}
void point_arry(){
    int array[4][5]={0};
    int i,j, k=0;
    for(int i=0;i<4;i++){
        for(j=0;j<5;j++){
            array[i][j]=k++;
        }
    }

    printf("*(array+1): %p\n",*(array+1));
    printf("array[1]:%p\n", array[1]);
    printf("&array[1][0]:%p\n",&array[1][0]);
      printf("**(array+1):%d\n",**(array+1));
      printf("array[1][0]:%d\n",array[1][0]);
      //array[i] ==  *(array+i)
      //*(*(array+i)+j) == array[i][j]

}
void point(){
    char a="1";
    int s=1;
    char *pa=&a;
    int *ps=&s;
    printf("a=%c\n", *pa);
    printf("s=%d\n",*ps);
    int num=5201314;
    int *p=&num;
    int **pp=&p;
    printf("%d\n",*p);
    printf("%d\n",**pp);
}
int count_str(){
   const char a[]="hello c";
    char *p=&a;
    int count=0;
    while (*p++ != '\0')
    {
       count++;
    }
    printf("count=%d\n",count);
    printf(" addr:=%p",a);
     printf(" addr:=%p",p);
    return count;
    
}
int main()
{
   signed int a = 10;
   unsigned short j=1;
   int arr[10]={0};
   int b[]={0};
  // char str2[10];
  // strcpy(str2,"hellow...zyy");
  // printf("%s\n",str2);
 //  strncpy(str2,"hellzyy",2);
 //  printf("%s\n", str2);

   char str[]="hellowrord";
   char str2[]="zyy";
//    strcat(str," ");
  // strcat(str,str2);
  str_cat();
   printf("str=%s\n",str);
    printf("str 的 sizeof=%d\n",sizeof(str));
    printf("str 的 lengeh=%d\n", strlen(str));

    printf("名字是%s,url：%s \n", NAME, URL);
    printf("int =%d\n",sizeof(int));
    printf("char =%d \n",sizeof(char));
    printf("short int=%d\n", sizeof(short));
    printf("long int=%d \n", sizeof(long));
    printf("long long int=%d \n", sizeof(long long));
    printf("_Boolen= %d \n", sizeof(_Bool));
    printf(" float = %d \n", sizeof(float));
    printf("double =%d \n ",sizeof(double));
    printf("long double=%d\n ",sizeof(long double));
    a++;
    printf("a=%d\n",a);
    printf("%d\n", 1 + (int)2.0);
    printf("%f\n", 1 + 2.0);
    int age;
    // printf("please input you age");
    // scanf("%d",&age);
    // if(age>10){
    //     printf(" age is %d",age);
    // }
    printf("%.0f\n", pow(2,10));
    print_mulit_table();

    point();
    count_str();
    point_arry();
    return 0;
}