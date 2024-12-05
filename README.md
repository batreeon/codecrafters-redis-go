[![progress-banner](https://backend.codecrafters.io/progress/redis/ea04b791-912f-4f6a-895f-a301812ee186)](https://app.codecrafters.io/users/codecrafters-bot?r=2qF)

This is a starting point for Go solutions to the
["Build Your Own Redis" Challenge](https://codecrafters.io/challenges/redis).

In this challenge, you'll build a toy Redis clone that's capable of handling
basic commands like `PING`, `SET` and `GET`. Along the way we'll learn about
event loops, the Redis protocol and more.

**Note**: If you're viewing this repo on GitHub, head over to
[codecrafters.io](https://codecrafters.io) to try the challenge.

# Passing the first stage

The entry point for your Redis implementation is in `app/server.go`. Study and
uncomment the relevant code, and push your changes to pass the first stage:

```sh
git commit -am "pass 1st stage" # any msg
git push origin master
```

That's all!

# Stage 2 & beyond

Note: This section is for stages 2 and beyond.

1. Ensure you have `go (1.19)` installed locally
1. Run `./your_program.sh` to run your Redis server, which is implemented in
   `app/server.go`.
1. Commit your changes and run `git push origin master` to submit your solution
   to CodeCrafters. Test output will be streamed to your terminal.

# Test

```sh
echo '*2\r\n$4\r\nECHO\r\n$9\r\nraspberry\r\n' | nc 0.0.0.0 6379
```

# 在Docker中运行redis

运行redis容器：

```sh
docker run --name <your_docker_name> -p 6379:6379 -d redis
```

redis客户端连接redis服务端，添加数据并生成rdb文件：

```sh
docker exec -it <your_docker_id or your_docker_name> redis-cli

// 在容器内终端运行：
set hello world
save
```

将容器中的rdb文件导出：

```sh
docker cp <your_docker_id or your_docker_name>:/data/dump.rdb <your_local_path>
```

查看rdb文件内容:

```sh
hexdump -C <your_local_path>/dump.rdb
```

文件内容输出：

```sh
00000000  52 45 44 49 53 30 30 31  32 fa 09 72 65 64 69 73  |REDIS0012..redis|
00000010  2d 76 65 72 05 37 2e 34  2e 31 fa 0a 72 65 64 69  |-ver.7.4.1..redi|
00000020  73 2d 62 69 74 73 c0 40  fa 05 63 74 69 6d 65 c2  |s-bits.@..ctime.|
00000030  d0 9a 51 67 fa 08 75 73  65 64 2d 6d 65 6d c2 c8  |..Qg..used-mem..|
00000040  eb 0f 00 fa 08 61 6f 66  2d 62 61 73 65 c0 00 fe  |.....aof-base...|
00000050  00 fb 01 00 00 05 68 65  6c 6c 6f 05 77 6f 72 6c  |......hello.worl|
00000060  64 ff 70 9e 7d 4b 33 0f  1d 80                    |d.p.}K3...|
0000006a
```
