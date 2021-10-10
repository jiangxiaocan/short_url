#短链接
1、直接运行go run main.go开启服务

2、访问http://127.0.0.1:8080/admin/create
参数url:获取得到的短链接<br>
`{
     "Status": 1,
     "Msg": "获取成功",
     "Data": {
         "Url": "http://www.baidu.com/20/x9os9l5zFE",
         "short_url": "03004eO3",
         "visit_url": "http://127.0.0.1:8080/03004eO3"
     }
 }`
 
 3、访问地址：http://127.0.0.1:8080/03004eO3 得到原来的地址
 <br>
 `{"Status":1,"Msg":"获取成功","Data":{"Url":"http://www.baidu.com/20/x9os9l5zFE","short_url":"","visit_url":""}}`
 
 4、创建短链接的身份校验还没有做，通过ip设置黑名单访问多少次失败拉入黑名单禁止再访问也没有做
 
 5、本项目基于gin框架，conf放配置文件
 
 6、有问题可留言