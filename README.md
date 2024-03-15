# Back-End
新苗同学后端

## database 
使用mysql作为持久化储存，redis作为缓存

看来看去，网上用于应用保持登录的方法无非就是session和jwt。
- jwt + redis
- jwt + redis + RSA（显然这两种方法有点违背了jwt设计的初衷）
- 长期token
- 双token jwt，用一个长期的Refresh Token 和 短期的Access Token（缺点是黑客获取了获取Refresh Token, 就可反复的刷新的Access Token）
- token以旧换新

jwt有个痛点就是由于他是无状态的，这样就会使得用户注销或者修改密码之后，原先的jwt在没有过期的时候仍然可用。这里有个解决方式就是将密钥设置为与用户相关的
