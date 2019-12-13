# 分别提交了两次

## 两次的debug信息
```
### 第一次还算正常
2019-12-13 17:05:53.199 [DEBU] [7 ms] UPDATE `bg_category` SET `parent_id`=0,`intro`='',`updated_at`=1576227953,`cover`='',`template`='',`id`=1,`cate_name`='teste',`slug`='test',`status`=1,`counts`=0,`list_order`=0,`created_at`=0 WHERE id=1
2019-12-13 17:05:53.200 [DEBU] [0 ms] SELECT COUNT(1) FROM `bg_category` 
2019-12-13 17:05:53.200 [DEBU] [0 ms] SELECT * FROM `bg_category` LIMIT 0,10
### 第二次就把slug作为条件了
slug,unique
2019-12-13 17:06:00.519 [DEBU] [12 ms] UPDATE `bg_category` SET `parent_id`=0,`intro`='',`created_at`=0,`updated_at`=1576227960,`template`='',`status`=1,`slug`='test',`counts`=0,`list_order`=0,`cover`='',`id`=1,`cate_name`='teste' WHERE slug='test'
2019-12-13 17:06:00.519 [DEBU] [0 ms] SELECT COUNT(1) FROM `bg_category` 
2019-12-13 17:06:00.521 [DEBU] [1 ms] SELECT * FROM `bg_category` LIMIT 0,10

```

POST信息

```
curl -X POST \
  http://127.0.0.1:8200/admin/post/category/edit/1 \
  -H 'Accept: */*' \
  -H 'Accept-Encoding: gzip, deflate' \
  -H 'Authorization: Bearer 154b683eca9d1a44d8fbc659bcac28f4' \
  -H 'Cache-Control: no-cache' \
  -H 'Connection: keep-alive' \
  -H 'Content-Length: 38' \
  -H 'Content-Type: application/x-www-form-urlencoded' \
  -H 'Host: 127.0.0.1:8200' \
  -H 'Postman-Token: 406af852-1e0e-4eff-8452-d2529b760037,9c9ba535-d913-4016-ab8b-48f392393f7f' \
  -H 'User-Agent: PostmanRuntime/7.20.1' \
  -H 'X-Requested-With: XMLHttpRequest' \
  -H 'cache-control: no-cache' \
  -d 'cateName=teste&slug=test&status=1&id=1'
```