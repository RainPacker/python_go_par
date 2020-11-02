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