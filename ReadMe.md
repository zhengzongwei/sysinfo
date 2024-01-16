# sysinfo

> 一个支持检索本机配置的命令行工具，支持 windows、linux、macos

## 如何使用

```bash
[root@VM-16-5-opencloudos work]# ./app_linux_amd64 -h
A command-line tool to query various system information

Usage:
  sysinfo [flags]
  sysinfo [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  cpu         Query CPU information
  disk        Query NET information
  help        Help about any command
  host        Query HOST information
  mem         Query memory information
  net         Query NET information

Flags:
  -h, --help   help for sysinfo

Use "sysinfo [command] --help" for more information about a command.

[root@VM-16-5-opencloudos work]# ./app_linux_amd64 cpu
+-------+--------------------------------+--------------+--------+-------+-------+----------+
|  CPU  |              NAME              |   VENDORID   | FAMILY | MODEL | CORES |   MHZ    |
+-------+--------------------------------+--------------+--------+-------+-------+----------+
| CPU 0 | Intel(R) Xeon(R) Gold 6133 CPU | GenuineIntel |      6 |    94 |     1 | 2499.998 |
|       | @ 2.50GHz                      |              |        |       |       |          |
| CPU 1 | Intel(R) Xeon(R) Gold 6133 CPU | GenuineIntel |      6 |    94 |     1 | 2499.998 |
|       | @ 2.50GHz                      |              |        |       |       |          |
+-------+--------------------------------+--------------+--------+-------+-------+----------+
[root@VM-16-5-opencloudos work]# ./app_linux_amd64 host
+---------------------+--------+------------+-------+-------+-------------+----------------+-----------------+------------+--------------------+----------------------+--------------------+--------------------------------------+
|      HOSTNAME       | UPTIME |  BOOTTIME  | PROCS |  OS   |  PLATFORM   | PLATFORMFAMILY | PLATFORMVERSION | KERNELARCH |   KERVERVERSION    | VIRTUALIZATIONSYSTEM | VIRTUALIZATIONROLE |                HOSTID                |
+---------------------+--------+------------+-------+-------+-------------+----------------+-----------------+------------+--------------------+----------------------+--------------------+--------------------------------------+
| VM-16-5-opencloudos |   7823 | 1705398012 |   119 | linux | opencloudos |                | 8.8.2305        | x86_64     | 5.4.119-20.0009.29 |                      | guest              | 56d1f483-49e5-44d8-bf07-5f327a1d2ec3 |
+---------------------+--------+------------+-------+-------+-------------+----------------+-----------------+------------+--------------------+----------------------+--------------------+--------------------------------------+
[root@VM-16-5-opencloudos work]# ./app_linux_amd64 mem
+------------+-----------+-----------+-------------+-----------+-----------+-----------+---------+
|   TOTAL    | AVAILABLE |   USED    | USEDPERCENT |   FREE    |  ACTIVE   | INACTIVE  |  WIRED  |
+------------+-----------+-----------+-------------+-----------+-----------+-----------+---------+
| 1721.50 MB | 969.06 MB | 566.47 MB | 32.91%      | 119.27 MB | 915.41 MB | 467.62 MB | 0.00 MB |
+------------+-----------+-----------+-------------+-----------+-----------+-----------+---------+
[root@VM-16-5-opencloudos work]# ./app_linux_amd64 net
+------+-------------------+-------+----------------------------+
| NAME |        MAC        |  MTU  |           ADDRS            |
+------+-------------------+-------+----------------------------+
| lo   |                   | 65536 | 127.0.0.1/8                |
|      |                   |       | ::1/128                    |
| eth0 | 52:54:00:92:22:0a |  1500 | 10.0.16.5/22               |
|      |                   |       | fe80::5054:ff:fe92:220a/64 |
+------+-------------------+-------+----------------------------+
[root@VM-16-5-opencloudos work]# ./app_linux_amd64 disk
+------+-----------+----------+----------+---------+-------------+
| PATH |  FSTYPE   |  TOTAL   |   FREE   |  USED   | USEDPERCENT |
+------+-----------+----------+----------+---------+-------------+
| /    | ext2/ext3 | 49.15 GB | 42.35 GB | 4.69 GB | 9.97%       |
+------+-----------+----------+----------+---------+-------------+


```


## 已支持

- [X] CPU
- [X] 网络
- [X] 磁盘
- [X] 内存
- [X] Host
