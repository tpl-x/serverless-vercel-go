# serverless-vercel-go
vercel serverless template for go
> for more detail,visit [vercel document](https://vercel.com/docs/functions/runtimes/go)

>vercel support creating kv and postgresql,after creation,you can get the connection info
> by lookup the env variables in the .env.sample

there are some limits of vercel go severless
+ cannot use the code under internal dir
+ the go.mod name should be like `github.com/tpl-x/serverless-vercel-go`
