# MacOs amd64 编译打包脚本

# 1.设置打包名称和格式方便修改
name="One-Encoder-Mac-amd64"
format="7z"

# 2.删除文件避免重复
rm -rf ./packed/mac/$name.$format
rm -rf ./build/One-Encoder
rm -rf ./build/One-Encoder.app

# 3.wails编译生成 ./build/One-Encoder
wails build -p

# 4.文件和附件移动到一个文件夹
mkdir ./packed/mac
mkdir ./packed/mac/$name
cp -R ./build/One-Encoder.app ./packed/mac/$name/
cp -R ./pack-assets/mac/ ./packed/mac/$name/

# 5.7z打包 && 删除临时文件夹
./pack-tools/mac/7z a -mm=LZMA2 -mx=9 ./packed/mac/$name.$format ./packed/mac/$name

# 6.提示
echo "打包成功！"

# 7.打开打包位置
open ./packed/mac

# 8.在包含附件/依赖的时候运行(mac无打包)程序
cp -R ./build/One-Encoder ./packed/mac/$name/
./packed/mac/$name/One-Encoder
# 9.删除打包测试用的临时文件夹
rm -rf ./packed/mac/$name

# 运行.app看不到终端执行fmt.Println
#open ./packed/mac/$name/One-Encoder.app
