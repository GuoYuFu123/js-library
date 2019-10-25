## class类

*created by 付果果   2019.10.25 19:19*

#### 定义：

```
在类中，我们通过new关键字来实例化对象，并将prototype指向类的prototype所指的对象，然而类构造函数外面通过点的语法定义的属性方法是不是添到新的实例化对象里去的，需要通过点的语法来调用!
```

#### 代码如下：

```
    var Book = function(id, bookName, price) {
        this.id = id;
        this.bookName = bookName;
        this.price = price;        
    }    
    Book.prototype.display = function () {
        console.log(this.id)
    }    
    //类的私有方法
    Book.getName = function() {
        console.log("getName")
    }
    var book = new Book(1, "js", 100);
    book.display();
    // book.getName(); //不能通过实例来访问

    // 通过点来定义的方法，只能通过点来调用，类似于java中的static语法
    Book.getName(); 
```

#### 注意

```
在类中，我们可以new多个实例，在类中通过this的数据都可以通过实例来获取到，但是，通过类名点方法名的函数或者属性，相当于类的静态的属性或者方法，只能通过类名点方法名或者变量名来进行调用
```

