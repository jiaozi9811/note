#(ERROR) err is shadowed during return
tags:golang

    func myTest() (retI int, retF float64) {
    	if true { //新的作用域
    		i0, retF := -1, -1.1
    		fmt.Println(i0, retF)
    		return //报错:[retF is shadowed during return]
    		//无返回值的return,会把返回值列表中的retI和retF返回出去,而不是新的作用域里的retF
    	}
    	return
    }
***正确写法***

    func myTest3() (retI int, retF float64) {
    	if true {
    		i0, retF := -1, -1.1
    		fmt.Println(i0, retF)
    		return retI, retF
    		//这里返回的是外层的retI和里层的retF,因为指明了,所以没有报错
    	}
    	return
    }
    
 分段详解

    func foo(x string) (ret int, err error) {//返回值列表定义了ret和err变量，作用域是整个函数体
    	if true {//新的语句块
    		ret, err := strconv.Atoi(x)  //这里又定义了新的变量ret和err，和返回值列表重名了。作用域是if语句块
    		if err != nil {
    			return  //这里的return语句会导致外层的ret和err被返回，而不是if语句里的ret和err
    		}
    	}
    	return ret, nil
    }

#**官方解释**
It's a new scope, so a naked return returns the outer err, not your inner err that was != nil.
So it's almost certainly not what you meant, hence the error.
<!--stackedit_data:
eyJoaXN0b3J5IjpbLTE2MTI5ODkzMzVdfQ==
-->
