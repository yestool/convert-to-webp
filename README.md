# convert-to-webp



## 启动容器

使用以下命令启动容器：

```bash
docker-compose up
```

如果你想在后台运行容器，可以加上`-d`参数：

```bash
docker-compose up -d
```

---

## 测试 API

容器启动后，你可以通过以下命令测试 API：

```bash
curl -X POST -F "image=@/path/to/your/image.jpg" http://localhost:8080/convert-to-webp --output converted.webp
```

---

## 停止和清理

### 停止容器

```bash
docker-compose down
```

### 删除镜像

如果你想删除构建的镜像，可以使用以下命令：

```bash
docker rmi <image-name>
```

---