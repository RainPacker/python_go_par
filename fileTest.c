#include <stdio.h>
#include <stdlib.h>
#include <time.h>
#define N 4
struct Stu
{
   char name[20];
   int numer;
   float score;
} stu[N],sb;

void stu_test(){
    FILE * fp;
    int i;
    if(( fp = fopen("score.txt","wb")) == NULL){
        printf("open file error");
         exit(EXIT_FAILURE);
    }
    printf("please input the stuend info(name no score):");
    for(i=0;i< N; i++){
        scanf("%s %d %f",stu[i].name, &stu[i].numer, &stu[i].score);
       
    }
    //写入
    fwrite(stu,sizeof(struct Stu),N,fp);
    fclose(fp);
     if((fp = fopen("score.txt","rb")) == NULL){
        printf("open file error");
         exit(EXIT_FAILURE);
    }
    fseek(fp,sizeof(struct Stu),SEEK_SET);
    fread(&sb,sizeof(struct Stu),1,fp);
    printf("sencond stu info, name=%s,number=%d,score=%f", sb.name,sb.numer, sb.score);
    fclose(fp);
}

int main(){
    //文件指针
    FILE *fp;
    struct tm *p;
    time_t t;
    time(&t);
    p = localtime(&t);
   printf("%d-%d-%d\n",1900+p->tm_year,1+p->tm_mon ,p->tm_mday);

      if((fp = fopen("hello.txt","w"))== NULL){
          printf("open file faile");

      }
      printf("%d\n",ftell(fp));
      fputs("F", fp);
       printf("%d\n",ftell(fp));
       fputs("ish", fp);
        printf("%d\n",ftell(fp));
     fclose(fp);

     stu_test();
    return 0;
}