class Watcher {
    constructor(vm, expr, cb) {
        this.vm =vm;
        this.expr =expr;
        this.cb = cb;
        this.oldVal = this.getOldVal();
    }
    getOldVal(){
            Dep.target = this;
            let oldVal = complieUtils.getValue(this.expr,this.vm);
            Dep.target = null;
            return oldVal;
    }
    //更新
    update(){
        let newVal =complieUtils.getValue(this.expr,this.vm);
        if(this.oldVal != newVal){
            this.cb(newVal);
        }
    }
}

class Dep{
    constructor() {
        this.subs =[];
    }
    //收集观察者
    addSub(watcher){
        this.subs.push(watcher);
    }
    //通知观察者
    notify(){
        console.info("通知观察者",this.subs)
        this.subs.forEach((w)=>{
                w.update();
        })
    }
}



class Observer {
     constructor(data) {
         this.observe(data);
     }
     observe(data){
         if(data&& typeof data === "object")
         {
             Object.keys(data).forEach((key)=>{
                    console.info("key..",key);
                    this.defindeReactive(data,key,data[key]);
             })
         }
     }
     defindeReactive(obj,key,value){
         //递归
         this.observe(value);
         const  dep =new Dep();
         Object.defineProperty(obj,key,{
             enumerable: true,
             configurable: true,
             get(){
                 //订阅者数据变化时 向Dep 中添加观察者
                 Dep.target && dep.addSub(Dep.target);
                 return value
             },
             set:(newVal)=>{
                 console.log(newVal)
                        this.observe(value);
                        if(newVal!==value){

                            value =newVal;
                        }
                   dep.notify();     
             }
         });
     }
}