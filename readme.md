# gopetstore

使用 go语言 实现的 jpetstore<br/>
不使用web框架进行整合，旨在上手go web编程<br/>
原 java 版: https://github.com/SwordHarry/Jpetstore<br/>

### 业务模块
- 商品模块
    - category
    - product
    - item
    - search
- 购物车模块
    - cart
- 用户模块
    - account
- 订单模块
    - order
    - lineItem
    - sequence

### 架构
template + go + mysql<br/>
**没有使用web框架，围绕 go 标准库http库**<br/>
采用 MVC 分层开发：DAO-persistence、service、controller、template<br/>
使用了 sessions 等第三方库<br/>

by the way，正在这里养成写注释和封装的习惯

![gopetstore](./gopetstore.png)