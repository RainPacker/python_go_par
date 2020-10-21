#include <iostream>
using namespace std;

int main(){
    int a=10;
    int b[10]={1,2,3,4,5,6,7,8,9,10};
    int *p = &a;
    int *p2=b;
    cout << p <<endl;
    cout <<*p2<<endl;
    cout<< sizeof(p)<<endl;
        cout<< sizeof(int)<<endl;
    cout <<"2234"<< endl;
    // cin >> a;
    cout <<&a<< endl ;
    return 0;
}