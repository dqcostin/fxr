# fxr
使用fscan联动Xray
## 简介
一款可以提取fscan扫描后的url，将url和xray进行联动扫描的工具，先提取fscan的http开头的url，然后进行去重，再利用xray进行批量扫描
## 用法
* 将fscan扫描完后的result.txt放在工具的同一级目录下
* 将xray_windows_amd64.exe工具也放在同级目录下
* 将xray的poc放在本地目录下（pocs的地址：https://github.com/chaitin/xray/tree/master/pocs）

![image](https://user-images.githubusercontent.com/82979945/127308592-bb4adcce-740c-4ab5-83b9-0b8b8781a72a.png)

## 运行截图

**在运行的同时，会产生一个test.txt文件，该文件是从fscan提取出来去重后的文件，以及运行结束会产生一个xray的html文件**

![image](https://user-images.githubusercontent.com/82979945/127308469-f28fb824-3497-4b6d-bf39-f9268cf6008d.png)

![image](https://user-images.githubusercontent.com/82979945/127308766-cab3c9fb-9f8b-48d4-bb5d-8658daf8e88a.png)
