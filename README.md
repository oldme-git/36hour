# 36hour
一个基于 `GoFrame` 的微服务空间管理项目，会集成各类云原生技术

# 当前技术栈
- GoFrame
- GRPC
- 服务发现：Etcd

# 部署方式
`CI/CD` 部署到 `k8s`。
当前流程：Github 工作流使用 Dockerfile 构建镜像上传到阿里云，rancher fleet 使用 kustomize 部署到 k8s。

# k8s服务端口分配
- 所有可用端口：32000 - 32700
- 32000: app/user 用户服务
- 32001: app/lib-manager 图书馆管理服务
- 32002: app/seat 座位管理服务