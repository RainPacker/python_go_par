const complieUtils={
    getValue(expr,vm){
       return  expr.split(".").reduce((data,current)=>{
            return data[current]
        },vm.$data)

    },
       text(node,expr,vm){
           let exprN;
           if(expr.indexOf("{{") !== -1){//{{person.name}}
               console.log(expr);
               exprN =expr.replace(/\{\{(.+?)\}\}/, (...args)=>{
                   console.log(args);

                   return args[1];
               })
               value=this.getValue(exprN, vm);
               this.updater.textUpdater(node,value)
           }else{
            value=this.getValue(expr, vm);
            this.updater.textUpdater(node,value)
           }
            
       },
       html(node,expr,vm){
        let value=this.getValue(expr, vm);
        this.updater.htmlUpdater(node,value)
       },
       model(node,expr,vm){
        let value=this.getValue(expr, vm);
        this.updater.modelUpdater(node,value)
       },
       on(node,expr,vm,eventName){
           let fn=vm.$options.methods && vm.$options.methods[expr] ;
         node.addEventListener(eventName,fn.bind(vm),false);
       },
       //更新
       updater:{
            modelUpdater(node,value){
                node.value=value;
            },
            textUpdater(node,value){
               
                    node.textContent=value;
            },
            htmlUpdater(node,value){
                node.innerHTML=value;
            }
       }
       

}
class Compile {
      constructor(el,vm){
          this.el= this.isElementNode(el)?el:document.querySelector(el);
          this.vm = vm;
          // 获取 文档碎片
          const fragment =this.node2fragment(this.el);
         console.log(fragment);
         //编译
         this.complie (fragment);

         this.el.appendChild(fragment);//添加文档到 页面中

      }
      isElementNode(node){
        return node.nodeType === 1;
      }
      node2fragment(el){
        const f= document.createDocumentFragment();
        let firstChild;
        while(firstChild = el.firstChild){
            f.appendChild(firstChild)
        }

        return f;

      }
      complie(fragment){
          const childNodes =fragment.childNodes;
          [...childNodes].forEach((child)=>{
              if(this.isElementNode(child)){
                 // console.log("元素节点",child)
                  //编译元素节点
                  this.complieElement(child);
              }else{
                 // console.log("文档节点",child)
                  //编译文本节点
                  this.complieText(child)
              }

              if(child.childNodes && child.childNodes.length>0){
                  this.complie(child);
              }
          })
      }
      /**
       * 是否是指令
       * @param {属性名称} name 属性名称 
       */
      isDirective(name){
          return name.startsWith("v-");
      }
      isEventName(name){
          return name.startsWith("@")
      }

      complieElement(node){
          //获取 标签上的属性
         var attributes = node.attributes;
         //console.log(attributes);
         [...attributes].forEach(attr=>{
        
         //获取属性名称，值
         const {name,value}= attr;
         console.log(name,value);
           if(this.isDirective(name)){//是一个指令 如 v-text,v-html,v-on:click
            
             const [,direct] = name.split("-");
             const [dirName,eventName]=direct.split(":");
             //更新数据 到视图
              complieUtils[dirName](node,value,this.vm, eventName);
             //删除所有指令标签上的属性
             node.removeAttribute("v-"+direct);   

           }else if(this.isEventName(name)){
                    let [,eventName]=name.split("@");
                    complieUtils["on"](node,value,this.vm, eventName);

           }


         })
      }
      /**
       * 编译文本
       * @param {}} node 
       */
      complieText(node){
         const  text=node.textContent;
         if(/\{\{(.+?)\}\}/.test(text)){
        //     console.info(text);
             complieUtils["text"](node,text,this.vm);
         }
      }


}


/**
 * 自己定义 vue 实现
 */
class MVue {

       constructor(options){
             this.$el=options.el;
             this.$data=options.data;
             this.$options=options;
             if(this.$el){
                 console.log("init.....")
                 //1实现数据观察者
                 //2.实现指令解析器
                 new Compile(this.$el,this);

             }
             
       }   
       

}



queryString = (url = '') => {
    const qs = {}
    url.replace(/([^?=&]+)(=([^&]*))?/g, ($0, $1, $2, $3) => {
        console.info($0,$1,$2,$3)
      if ($3 === undefined) {
        return
      }
      try {
        qs[$1] = decodeURIComponent($3)
      } catch (e) {
        qs[$1] = $3
      }
    })
    return qs
  }

  console.info(queryString("www.badi.com?aa=b&&c=799&9=1"))

  let reg2=/[123456]/
  console.log(reg2.test("7"))
let reg3=/(12|23)/
console.log(reg3.test("234"));

let reg4=/\d/g;
console.log("123456".match(reg4));

let  str2="200353453Wd_-&4";
console.log("=============")
console.log(str2.match(/\d/));
console.log(str2.match(/\d+/));
console.log(str2.match(/\w/));
console.log(str2.match(/\w+/g));
console.log(str2.match(/\W/));
console.log(str2.match(/\W+/g));
console.log(str2.match(/\D/g));
console.log(str2.match(/\D+/));
console.log(str2.match(/\D+/g));

console.log("=============")
let str3=`432432423rewre@^^^454
  werwer4w3555*
`
console.log(str3.match(/./));
console.log(str3.match(/./g));
console.log(str3.match(/.+/));
console.log(str3.match(/.+/g));
//忽略换行
console.log(str3.match(/.+/s));
console.log("=============")
let str4=`
 <span>
    12423jowieroruoui4u334234jsdj@*@*
    </span>
`;
console.log(str4.match(/<span>[\s\S]+<\/span>/))
console.log(str4.match(/\s/g))
console.log(str4.match(/\s+/))
console.log(str4.match(/\s+/g))
console.log("===========")
let str5='UUuUNNNnnnnu';
console.log(str5.match(/u/i));
console.log(str5.match(/u/gi));
console.log(">>>>>>>>>>>>>>>>")
let str6=`
   #1 js, 200 # 
   #2 php, 210 #
   #3 react, 500 # 哈哈
   #4 vue, 100 #
   #5 python, 90 #
   #6 c++, 100 #
   #7 java, 300 #
   #8 c, 300 #
`

//匹配所有行 获取 name ,价格 并返回为 obj
let arrs=str6.match(/^\s*#\d+\s+.+\s+#$/gm);
console.log(arrs)
arrs=arrs.map((item=>{
   v= item.replace(/\s*#\d+\s*/,"").replace(/\s+#/,"");
   console.log(v);
   let [name,price]= v.split(",");
   console.log(name,price)
   name=name.replace(/\s*/,"");
   price=price.replace(/\s*/,"");
   return {name,price};
}))
console.log(JSON.stringify(arrs,null,2))
console.log("+++++++++++++++")

let srtr7 ='sabst.erw, 加油打工人';
//
console.log(srtr7.match(/\p{L}/gu));
//逗号
console.log(srtr7.match(/\p{P}/gu));
//汉字
console.log(srtr7.match(/\p{sc=Han}+/gu));
console.log(srtr7.match(/\p{sc=Han}/gu));

console.log("***********")
let str8= "jiayoudagongren...";

console.log(str8.match(/\w/))
console.log(str8.match(/\w/g))
let reg8= /\w/g ;
//console.log(reg8.exec(str8));

while ((res= reg8.exec(str8))){
    console.log(reg8.lastIndex, res);
}



