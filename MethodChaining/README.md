## 递归,链式调用,你真的需要这些吗?



## 链式调用会影响性能吗?

`链式调用` : .当一个对象的方法返回值是对象本身时,可以继续调用其他方法,这样就叫做链式调用.

```js
ladder.up()
    .up()
    .down()
    .showStep()
    .down()
    .showStep()
```

当然了你可以将其写成这样的形式.

```js
ladder = ladder.up()
ladder = ladder.up()
ladder = ladder.down()
ladder = ladder.showStep()
ladder = ladder.down()
ladder = ladder.showStep()
```

或者你可以保留每个时刻的状态.

```
ladder_status1 = ladder.up()
ladder_status2 = ladder_status1.up()
ladder_status3 = ladder_status2.down()
ladder_status4 = ladder_status3.showStep()
ladder_status5 = ladder_status4.down()
ladder_status6 = ladder_status5.showStep()
```

## 实验设计

这里其实应该就是对不同的语言设计实验的就是说

