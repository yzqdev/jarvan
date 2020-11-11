[![Version](https://img.shields.io/badge/version-1.0.0-green.svg)](https://jarvan)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](http://opensource.org/licenses/MIT)


#### AP列表

- /api/auth/login 登录/身份校验
- /api/auth/logout 登出
- /api/article/list 文章列表
- /api/article/detail/:id 获取文章详情
- /upload 图片上传
- /upload/images 获取图片链接

完善中...

#### 数据结构

|字段|描述|
|-----|------|
|code|返回状态码 200-success|
|data|返回数据|
|msg|返回消息|

eg:
```json
{
    "code": 200,
    "data": {
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InRhbmdjaHVubGluaXRAZ21haWwuY29tIiwicGFzc3dvcmQiOiJtYW50aXMxMjEyIiwiZXhwIjoxNTQ3MzU4OTkwfQ.Wff2gVySzR7ARCa1BAhGLRBnOhsw1Y4R9RC4aIkGEMw",
        "username": "Mantis"
    },
    "msg": "操作成功"
}
```
