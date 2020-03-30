#指定基础镜像
FROM centos
#指定镜像元数据
LABEL auto=OnlyOneFace
#指定端口 网络协议
EXPOSE 8081
#运行的程序
ENTRYPOINT exec /search/service/DistributedCrawler/DistributedCrawler