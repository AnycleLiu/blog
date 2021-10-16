# create
curl -X POST "http://localhost:8080/v1/article" -H "content-type:application/json" -d '{"title":"golang channel原理","content":"channel实际上是hchan这个结构，对于有缓冲channel，内部有一个ring buf","tag":["golang","源码"]}' -i

# get
 curl http://localhost:8080/v1/article/1