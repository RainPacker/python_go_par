#include <stdio.h>
#include <math.h>
#include <string.h>
#include <stdarg.h>
#include <stdlib.h>
#define URL "http://www.fishc.com"
#define NAME "杨杨"
#define MAX(x,y) ((x)>(y)?(x):(y))
struct Book addBook(void);
int sum(int,...);
char *getWord(char c);
void mem();
void hannu(int ,char ,char , char );
int calulate(int (*)(int,...),int, ...);
void nodeTest(void);
struct Book
{
  char title[28];
  char name[40];
  float price;
};

struct Book addBook(){
    struct Book book;
    printf("请输入书名：");
    scanf("%s",book.name);
     printf("请输入标题：");
     scanf("%s",book.title);
      printf("请输入价格：");
      scanf("%f",&book.price);
printf("book.name=%s,title=%s,price=%.2f", book.name, book.title, book.price);
    return book;

}

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

int sum(int n,...){
    int i ,sum=0;
    va_list vap;
    va_start(vap,n);
    for(i=0;i<n ;i++){
        sum+=va_arg(vap, int);
    }
    va_end(vap);
    return sum;
}


char *getWord(char c){
    switch (c)
    {
    case 'A':
       return "Apple";
    case 'B':
        return "Banana";
    default:
       return "None";
    }
}

int calulate(int (*fpsum)(int,...),int n,...){
     va_list args;
     va_start(args,n); 
     int result=(*fpsum)(n,args);
     va_end(args);      //初始化参数指针，将ap指向第一个实际参数的地址
    return  result;
}


void hannu(int n,char x,char y, char z){
    if(n==1){
        printf("%c-->%c\n",x,z);
    }else{
        hannu(n-1,x,z,y);
        printf("%c-->%c\n",x,z);
        hannu(n-1,y,x,z);

    }
}
void mem(){
     int *ptr;
     ptr=malloc(sizeof(int));
     *ptr = 11;
     //free
     printf("mem:ptr:%d", *ptr);
     printf("mem addr%p", ptr);
     free(ptr);
}
typedef struct Node
{
 int value;
 struct Node *next  ;
}NODE,*PNODE;
void insetNode(struct Node **head ,int value){
          struct Node *previous;
          struct Node *current;
          NODE *new;
          current = *head;
          previous = NULL;
          //找到了比 value 小的地方插入数据 
          while (current !=NULL && current->value < value){
                    previous = current;
                    current = current->next;
          }
           //分配空间
           new = (struct Node *) malloc(sizeof(struct Node));
           if(new == NULL){
               printf("malloc memory error!\n");
           }
           new->value = value;
           new->next = current;
           if(previous == NULL){
               *head = new;

           }else{
               previous->next = new;
           }
          
}

void printNode(struct Node *head){
    struct Node *current;
    current = head;
    while (current != NULL)
    {
       printf("%d ",current->value);
       current = current->next;
    }
    putchar('\n');
    
    

}
void delNode(struct Node **head,int value){
    struct Node *current;
    struct Node *previous;
    current = *head;
    previous = NULL;
    while (current != NULL && current->value != value)
    {
            previous = current;
            current = current->next;
    }
    if (current == NULL){//未找到，或者是头节点
        printf(" can not find the match value in the node \n");
        return ;
    }else{
          if(previous == NULL){
                *head = current->next;

          }else
          {
            previous->next = current->next;
          }
          free(current);
          

    }
    
}
//链表测试
void nodeAddTest(void){
   struct Node *head = NULL;
   int input;
   while (1)
   {
      printf(" plese input a numer(-1 for exit):");
      scanf("%d",&input);
      if(input == -1){
          break;
      }
      insetNode(&head, input);
      printNode(head);
   }
   

}

void nodeDelTest(){
     struct Node *head = NULL;
   int input;
   while (1)
   {
      printf(" plese input a numer(-1 for exit):");
      scanf("%d",&input);
      if(input == -1){
          break;
      }
      insetNode(&head, input);
      printNode(head);
   }
   
    while (1)
   {
      printf(" plese input a numer to del(-1 for exit):");
      scanf("%d",&input);
      if(input == -1){
          break;
      }
      delNode(&head, input);
      printNode(head);
   }




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
    int result=sum(3,1,2,3);
    printf("sum(3,1,2,3)=%d\n",result);

    // char s2[]=getChar('A');
    printf("%s\n",getWord('A'));
    char *(*fp)(char);
    int (*fsum)(int,...);
    fsum=sum;
    fp=getWord;
     printf("%s\n",(*fp)('B'));
      printf("sum(3,3,4,5)=%d\n",(*fsum)(3,3,4,5));
       int res= calulate(sum,3,1,2,3);
        printf("sum(3,1,2,3)=%d\n",res);
        hannu(6,'x','y','z');

        //
       // addBook();
       struct Book book={
           .name="带你飞。。",
           .title="zzzzz"
       };
       struct Book *ptBook;
       ptBook=&book;
       book.price=1.23;
       //book.name="zyy";
      // book.title="23423";
      printf("book.name=%s\n",ptBook->name);
      printf("book.title=%s\n",(*ptBook).title);
      mem();
      printf("\n");
    // nodeTest();
    nodeDelTest();
      
    return 0;
}