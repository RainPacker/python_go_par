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