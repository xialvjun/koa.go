middleware 赋值到 request 上的属性有两个地方可以存储：
1. 存储在 middleware 上，如 session 是包名，store := session.NewStore(), store 内部有个 store.map[*http.Request]Session，store 在下面所有的 middleware/handler 中可用。注意：这里的 map 应该被 mutex 包裹起来，而 key 也是 *http.Request，因为这里的 store 说的是在一个请求中的上下文传递，而 session.GetRemote(request.Cookie.sid) 则是从数据库的 session 表中取，session.Get(request) 这是从内存中读取
2. 存储在 req 对象上，可以参照 volatile 库，也可以参照 http://blog.csdn.net/u014029783/article/details/53782864 放在 request.Context 上。这里把数据放在 req 对象上，要给它属性名（map 的 key），属性名约定为 package name，即 map[string]interface{} { "github.com/xialvjun/koa.go.session": SessionObject }，或者属性名由使用者决定。之后由 package 提供方法，session.Get(request)：这里是读取 request 上的 package_name 属性，如果属性名由使用者决定，则 package 内有个全局变量，或者新生成一个结构体，name 放在结构体上

取出数据后都用类型断言强转，返回断言值和断言ok。。。 result, ok := value.(Session); return result, ok

a := context.WithValue(context.Background(), "a", 123)
b := context.WithValue(a, "b", []int{1, 2, 3})
if t, ok := b.Value("b").([]int); ok {
  fmt.Print(t)
}
fmt.Print(b.Value("b"))


middleware 不需要 ResponseWriter...应该是所有的 middleware 都返回数据。。。然后上层 midmiddleware 包装下层 middleware 返回的数据，最终由框架来 Write
