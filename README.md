# Egin

## 总结
封装框架的优点是可以减少重复代码的使用，将粒度极细的接口隐藏，从更宏观的视角调用

针对使用场景，封装*http.Request和http.ResponseWriter的方法，简化相关接口的调用，是设计 Context 的原因之一

## EveryDay Feature
> Day1: 实现了静态路由，支持解析GET，POST请求并调用对应函数，建立了Egin的框架，50行

> Day2: 将路由(router)独立出来，方便之后增强。设计上下文(Context)，封装 Request 和 Response ，提供对 JSON、HTML 等返回类型的支持。动手写 EGin 框架的第二天，框架代码140行，新增代码约90行

## 2024/3/7
建库

## 2023/3/8 
Accept Day1, Day2