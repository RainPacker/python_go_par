
window.onload=function (){
  var canvas =document.getElementById("canvas");
 var ctx= canvas.getContext("2d");
 ctx.lineWight= 3;//线条的粗
 ctx.strokeStyle = 'red'
//  ctx.moveTo(10,10);
//  ctx.lineTo(100,10);
//  ctx.lineTo(100,200);
//  ctx.stroke();
//绘制矩形
ctx.rect(0,0,300,150);//ctx.rect(在上角 x,左上角 y,width,height)
ctx.fillStyle= "#ccc"//填充颜色
ctx.fill()//开始填充
//ctx.clearRect(50,50,50,50)//清除区域（））
canvas.onmousedown= function(e){
    canvas.onmousemove =function(e){
        ctx.clearRect(e.clientX,e.clientY,10,10)//清除区域（））
    }
   
}
canvas.onmouseup =function(){
    canvas.onmousemove = null; //清除事件
}
var arr=["一个亿","一等将","二等将","谢谢惠顾"]
var prise= document.querySelector(".prize");
var i =Math.floor(Math.random()*arr.length);
console.info(i)
prise.innerHTML=arr[i];


//  canvas.onmousemove= function(e){
//      console.info(e.clientX,e.clientY);
//      ctx.lineTo(e.clientX,e.clientY);
//      ctx.stroke();
//  }
//  canvas.ontouchmove=function(e){
//      var touch=e.changedTouches[0]
//      console.info(touch)
//     console.info(touch.clientX,touch.clientY);
//     ctx.lineTo(touch.clientX,touch.clientY);
//     ctx.stroke();
// }
  
}

var obj={
    "name":"zyy",
    "age":10
}
var obj2={};
//原型指向 obj 
obj2 = Object.create(obj,{
     sex:{
         value:"M",
         writable: true,
         configurable: true,
         enumerable: true
         
     }
});
console.log(obj2)

Object.defineProperties(obj2,{
    fullName:{
        get:function(){
            console.info("get....")
            return "!!!!!|";
        },
        set:function(newVal){
            console.info("set....",newVal)
            fullName= newVal
        }

    }
})
console.log(obj2);

function callDemo(data){
    console.info(this,data)
}
callDemo(obj)
callDemo.apply(obj,[33]);//第二个必须传入数组为参数
callDemo.call(obj,33);//直接从第二个参数依次传入参数
//bind绑定不会立即执行，而是返回方法
var bar=callDemo.bind(obj);
callDemo.bind(obj)(3333)
console.info(bar);
bar(obj);
//sample
setTimeout(function(){
   console.log(this);
}.bind(obj),1000);
//es 5 array
let arr =[2,3,5,7,8,0];
console.log(arr.join("+"));
//console.info(arr.toSource());
arr.findIndex((val,index,obj)=>{
    console.info(val,index,obj)
})


let a = 1;
let ob = {
a:2,
sayHi:()=>{
  console.log("sayHai");
console.log(this,this.a)
}
}
ob.sayHi();

//js 正则
let str = "abccke20000yyy999";
//使用 es6 查询所有字符 
 let num= [...str].filter(a=> !Number.isNaN(parseInt(a))).join("")
console.log("num=>",num);
//使用正则查询 
console.log(str.match(/\d/g).join(""))

//正则字变量
let hd="hdstr.com";
let testStr= 'hd';
console.log(/hd/.test(hd));
console.log(eval(`/${testStr}/`).test(hd))
//使用对象创建正则
let reg= new RegExp(testStr,"g");
console.log(reg.test(hd))
//区间匹配
console.log("#########################")
let str6 = "uuuuzzzuuaaadd 4234zzzz";
console.info(str6.match(/[ud]/i))
console.info(str6.match(/[ud]/gi))
console.info(str6.match(/[ud]+/gi))
console.info(str6.match(/[^ud]/))
console.info(str6.match(/[^ud]+/))
let str7 =`
     <h1>11</h1>
     <h3></h3>
     <h2>hhherwereeeeerwer</h2>
     <div></div>
`

console.log(str7.match(/<h([0-6])>([\s\S]*)<\/h\1>/))
console.log(str7.match(/<h([0-6])>([\s\S]*)<\/h\1>/g))

str7 = str7.replace(/<(h[1-6])>([\s\S]*)<\/\1>/ig, "<p>$2</p>");
console.log(str7);
//域名检测
let url="https://baid.com";

let reg7 =/(?<pro>https?):\/\/((?:\w+\.)*\w+\.(?:com|org|net|cn))/i
console.log(url.match(reg7)) 
console.log(reg7.exec(url));
//批量验证 密码 ：只能是数字字母  要有大定字母
let pwd =`abc12345gG`;
var regs8=[
    /^[a-z0-9]{5,10}$/i,
    /[A-Z]/
];
regs8.every(e=>{
    console.log(e);
    console.log( e.test(pwd))
    return  e.test(pwd);
   
})
console.log(pwd.match(/[A-Z]/))

//禁用贪婪
let str9 =`hdddd`;
console.log(str9.match(/hd+?/))
console.log(str9.match(/hd*?/))
console.log(str9.match(/hd{0,1}?/))
console.log(str9.match(/hd{1,4}?/))


//mathAll 全局匹配获取标签中的内容
let  htmlStr= `
    <h2>qewqwe</h2>
    <h1>qweqw</h1>
    <h3>www</h3>
    <h4>1weeqw
    </h4>
`;
let reg10 = /<(h[1-6])>([\s\S]+?)<\/\1>/ig;
let reg11 = /<(h[1-6])>([\s\S]+?)<\/\1>/i;
//使用 mathch 获取不到标签中的内容
console.log(htmlStr.match(reg10));
//此时使用 mathcAll
//console.log(let  v= htmlStr.matchAll(reg10));
let  v= htmlStr.matchAll(reg10)
let data_target =[];
for (const item of   v) {
     console.info(item[2])
     data_target.push(item[2]);
}
console.info(data_target)

// matchAll 兼容写法
String.prototype.matchAll=function(reg){
    let res = this.match(reg);
    if(res){
        let str = this.replace(res[0],'^'.repeat(res[0].length));
        console.info("str>>",str)
        let match = str.matchAll(reg)||[]
        return [res,...match];
    }

}

let  v2= htmlStr.matchAll(reg10)

//使用 正则 exec 方法来实现全局匹配

while(res =reg10.exec(htmlStr)){
    console.info("res>>", res[2]);
    
}

let str11= "2020/11/12";
let p2= "2020-11-12";
let reg12= /[-/]/;

 console.log(str11.split(reg12)); 
console.log(p2.split(reg12))

//replace 替换使用正则
let date= '2020/11/12';

 console.log(date.replace(/\//g,"-"));
 let tel ="(010)9999999 (020)8888888";
 let reg13 =/\((?<preffix>\d{3,4})\)(\d{7,8})/g
 console.log(tel.match(reg13));
 // $`（匹配元素在）  $&(匹配元素本身相) $'（匹配元素右）
 console.log("",tel.replace(reg13, "$1-$2"));
 console.log("$&",tel.replace(reg13, "$&"));
 console.log("$`", tel.replace(reg13, "=$`"));
 console.log("$'", tel.replace(reg13, "$'"));
let tel2 ='18051931895';
console.log(tel2.replace(/\d{1,11}/g,"<tel>$&</tel>"))

let htmlStr14= `
     <a href="www.badi.com">a</a>
     <a href="www.xm.com">b</a>
     <a id="xx" href="www.qq.com">c</a>
`;
//匹配获取 所有的连接和名称

const reg14 =/<a.*?href=(["'])(?<link>.*?)\1>(?<title>.*?)<\/a>/i


console.log(htmlStr14.match(reg14));
console.log(htmlStr14.matchAll(reg14));
let infofs =[];
htmlStr14.matchAll(reg14).forEach(element => {

    console.info(">>>>>>>>>>",element.groups)
    infofs.push(element["groups"]);
});
console.log(infofs)

//断言
let str15 = "zhangyangyang zhangyangyang123";
//使用（?=条件）表示条件要出现的 (?!)
//表示只匹配符合条件的 如下表示只匹配符合后面有123的「zhangyangyang」
console.log(str15.match(/zhangyangyang(?=123)/ig))

let lessons = `
    js, 200元,300次
    php, 300.00元,200次
    pythony,100元,100次 
`;
//将对出现的金额进行补零,使用断言限制条件
let reg16 = /(\d+)(\.00)?(?=元)/ig

lessons=lessons.replace(reg16,(v,...args)=>{
      console.info(args);
      args[1] = args[1]||".00"; 
      return args.splice(0,2).join("");
});
console.info(lessons);
let users =`
    18051931895 17156311991
    17156311991
    17156311991
`;
//进行号码模糊处理
let reg17 =/(?<=\d{7})\d{4}/ig
users =users.replace(reg17,(...args)=>{
    return "*".repeat(4);
})
console.info(users)

let str18 ='1312312zhaNg21313131';
//断言字符不能有 zhang
let reg18 = /^(?!.*zhang.*).*/i
console.log(reg18.test(str18));
console.log(str18.match(reg18));

//获取所有域名除了含有 qq
let urls =`
        <a href="https://www.baidu.com/1"></a>
        <a href="https://oss.baidu.com/2"></a>
        <a href="https://qq.com/3"></a>
`;
let reg19 =/https?:\/\/(\w+)?(?<!qq)\..+?(?=\/)/ig ;
console.log(urls.match(reg19));

//symbol 的使用
let sy =Symbol("zzzzz");
console.info(typeof sy);
console.info(sy.description);
let cmf =Symbol.for("zyy");
console.log(Symbol.keyFor(cmf));

let obj_sym= {
    [sy]:"symbol 属性",
    name:"zyy"
}
console.log(obj_sym[sy]);//key 为 Symbol 类型的 key  相当于私有
console.log("_______________________")
//无法获取 Symbol 类型的属性
for (const key in obj_sym) {
    console.info(key);
   
}
console.log("_______________________")
//无法获取 Symbol 类型的属性
for (const iterator of Object.keys(obj_sym)) {
    console.info(iterator)
}
//使用Reflect.ownKeys 获取所用 key 包括 Symbol 类型的 key
console.log("_______________________")
for (const iterator of Reflect.ownKeys(obj_sym)) {
        console.info(iterator);
}

let user1 = {
    name:"李四",
    key :Symbol()
}
let user2 = {
    name:"李四",
    key: Symbol()
}

let grade = {
    [user1.key]:{
        js:90,
        java:100
    },
    [user2.key]:{
        js:90,
        java:88
    }
}
console.log(grade)
console.log(grade[user1.key])
console.log(grade[user2.key])
console.log("_______________________")
let [name, ...args] = ["zyy","zhang","yangyang"];

console.log(args);
console.log(...args);

let arr_1 = [111,222,333,444];
//arr_1.push(22,33);//从后面插入
arr_1.unshift(0);//从前面插入一个新值

arr_1.pop();
arr_1.shift();
console.table(arr_1);

let arr_2 = [111,222,2,444];
function array_move(array,from, to){
    if(from<0|| to> array.length){
        console.error("参数错误");
        return ;
    }
    let arr_new = [...array];
      let item = arr_new.splice(from,1);
      arr_new.splice(to, 0, ...item);
      return arr_new;
}

console.log(arr_2.includes(111));
console.table(array_move(arr_2,1,3));
arr_2.findIndex(item=>{
    return item ==222
})
arr_2.find(item=>{
    return item ==222
})

arr_2.sort((a,b)=>{
    //-1 从小到大 1 从大到小
    return a-b;
})
//every() 方法测试一个数组内的所有元素是否都能通过某个指定函数的测试。它返回一个布尔
let re1 = arr_2.every(item=>{
    return item>2;
})
//some() 方法测试数组中是不是至少有1个元素通过了被提供的函数测试。它返回的是一个
let re2 = arr_2.some(item=>{
    return item>2;
})
console.log(re1, re2)
console.log(arr_2.entries())
for (const item of arr_2.entries()) {
    console.log(item)
    
}
let res3 = arr_2.filter((item)=>{
    return item>10;
})
console.log(res3)
//map() 方法创建一个新数组，其结果是该数组中的每个元素是调用一次提供的函数后的返回值。
let res4 = arr_2.map((item,index,arr)=>{
    console.log(item, index, arr);
    return item+1;
})
console.log(res4)

     let obj_t = {a:"zz",b:11};
     let t2={
         a:"yyy",
         ...obj_t
     }
     console.log(t2)
console.log("_______________________")
//reduce() 方法对数组中的每个元素执行一个由您提供的reducer函数(升序执行)，将其结果汇总为单个返回值。
let res5 = arr_2.reduce((pre,cur,index,arr)=>{
    console.log(pre,cur,index);
    return pre+cur;
})
console.log(res5)
//返回数组中最大的值
let res6=arr_2.reduce((pre,vale)=>{
    return pre>vale?pre:vale
})
console.log("maxNumer",res6)  

let cart = [
    {name:"js",price:100},
    {name:"xm",price:190},
    {name:"xm",price:190},
    {name:"huawei",price:290},
    {name:"huawei",price:290},
    {name:"iphone",price:1000},
    {name:"pythos",price:20},
 

]
//查询数据cart中 中价格最贵的商品
let res7 = cart.reduce((pre,item,index)=>{
     if(pre.price > item.price){
         return pre;
     }else{
         return item;
     }
})
//获取总价
let res8 = cart.reduce((total,item)=>{
    return total+ item.price;
},0)
console.log("totalPrice",res8);
//获取金额大于200的所有商品名称
let pro = cart.reduce((arr,item)=>{
    if(item.price > 200 ){
        arr.push(item);
    }
    return arr;
},[]).map(item=>item.name)
console.table(pro)

//去除重复的商品
let pro_dist = cart.reduce((arr,cur)=>{
      let find =arr.find(item=>{
          if(item.name == cur.name){
              return item;
          }
      });
      if(!find) arr.push(cur)
      return arr;
       
},[])
console.table(pro_dist)