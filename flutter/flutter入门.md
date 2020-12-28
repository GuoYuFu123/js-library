### Flutter入门

*author： 付国宇    date：2020.11.29*

### 环境配置

##### 一、获取Flutter SDK

###### 1、去flutter官网下载其最新可用的安装包，[转到下载页](https://flutter.io/sdk-archive/#macos)

###### 2、将flutterSDK进行解压【最好放到一个指定的目录下，后面设置path路径还需要用】

###### 3、添加环境变量

```bash
❯ cd ~
❯ vim .zshrc

# 将下面的path拷贝到文件最后， esc => :wq! 保存即可
# Flutter Path
export PATH="$PATH:[pwd]/flutter/bin"
# eg: export PATH="$PATH:/Users/fuguoyu/flutter/sdk/flutter/bin"

❯ source .zshrc
```

4、可以通过命令来看一下是否配置成功

```bash
❯ flutter  
# 或者
❯ flutter -h
```



### 二、检查是否依赖完成

运行以下命令查看是否需要安装其它依赖项来完成安装，根据错误信息进行操作就可以

```bash
❯ flutter doctor
```

输出

```bash
Doctor summary (to see all details, run flutter doctor -v):
[✓] Flutter (Channel stable, 1.22.4, on Mac OS X 10.15.6 19G73 darwin-x64, locale zh-Hans-CN)
[✗] Android toolchain - develop for Android devices
    ✗ Unable to locate Android SDK.
      Install Android Studio from: https://developer.android.com/studio/index.html
      On first launch it will assist you in installing the Android SDK components.
      (or visit https://flutter.dev/docs/get-started/install/macos#android-setup for detailed instructions).
      If the Android SDK has been installed to a custom location, set ANDROID_SDK_ROOT to that location.
      You may also want to add it to your PATH environment variable.

[!] Xcode - develop for iOS and macOS (Xcode 12.2)
    ✗ CocoaPods not installed.
        CocoaPods is used to retrieve the iOS and macOS platform side's plugin code that responds to your
        plugin usage on the Dart side.
        Without CocoaPods, plugins will not work on iOS or macOS.
        For more info, see https://flutter.dev/platform-plugins
      To install:
        sudo gem install cocoapods
[!] Android Studio (not installed)
[✓] VS Code (version 1.51.1)

[✓] Connected device (1 available)

! Doctor found issues in 3 categories.
```

因为咱们现在都是mac，模拟器我们用xcode就可以了

##### （一）、安装xcode

###### 1、这个去app store就可以去下载，最好在家下载，或者网络好的地方下载，我在公司下载了用了8个小时左右

###### 2、配置Xcode命令行工具以使用新安装的Xcode版本

```bash
sudo xcode-select --switch /Applications/Xcode.app/Contents/Developer
```

###### 3、确保Xcode许可协议是通过打开一次Xcode或通过命令，如果有[Y/N]的话，一路Y就可以了，N的话会有问题

```bash
sudo xcodebuild -license
```

###### 4、安装包管理工具

```bash
# 通过brew安装(推荐)
brew install cocoapods
brew link --overwrite cocoapods

# 或者通过gem安装
sudo gem install cocoapods

# https://github.com/orta/cocoapods-keys/issues/198#issuecomment-510909030
# https://github.com/CocoaPods/CocoaPods/issues/9602
```

##### （二）、配置vsCode

###### 在扩展中安装Flutter插件，并重新启动加载插件

##### （三）设置ios模拟器

```bash
# 打开一个模拟器，确保您的模拟器正在使用64位设备（iPhone 5s或更高版本）
open -a Simulator
```



### 三、android studio

##### 1、[官网下载IDE](https://developer.android.com/studio?gclid=CjwKCAiA_eb-BRB2EiwAGBnXXjMc5XGfiZOBDtCQB9LWfS4_DKYMdoeWHMtRvknWDeRmNQPwpwM2SRoCv_wQAvD_BwE&gclsrc=aw.ds)

##### 2、打开IDE，安装flutter插件和dart插件

```
Android Studio -> preperences -> plugins (右侧) markplace 搜索flutter安装即可，会自动提示安装Dart
```

##### 3、但是运行flutter doctor 会一直显示plugins没有安装上

```
1、https://github.com/flutter/flutter-intellij/issues/4523
2、https://stackoverflow.com/questions/64395106/update-to-android-studio-4-1-flutter-plugin-and-dart-plugin-not-installed/64417280#64417280

reason：对于android studio文件的路径变了，目前flutter sdk的beta版已修复
answer：1）软链 
				2）切换flutter sdk为beta版
        				flutter channel dev
								flutter upgrade
```

##### 4、打开Android Studio，新建虚拟机， 第一次会去谷歌下载gradle，这个东西很容下载失败，如果失败，可以多反复试几次

##### 5、下面就和mac的使用方式一样了....

<----------👆以上环境就配置好了👆--------->

### 初始化项目

### 一、创建

###### 1、打开vscode，创建myapp项目

```bash
cmd + shift + p
=> Flutter: New Project
=> ...
```

2、连接虚拟机，或者模拟器

```
在vscode的下方会有一个No device，点击连接设备
```

###### 3、运行myapp

```bash
❯ cd myapp
❯ flutter run
Launching lib/main.dart on iPhone 12 Pro Max in debug mode...

Running Xcode build...
 └─Compiling, linking and signing...                         5.3s
Xcode build done.                                           18.5s
Waiting for iPhone 12 Pro Max to report its views...                 4ms
Syncing files to device iPhone 12 Pro Max...                       176ms

Flutter run key commands.
r Hot reload. 🔥🔥🔥
R Hot restart.
h Repeat this help message.
d Detach (terminate "flutter run" but leave application running).
c Clear the screen
q Quit (terminate the application on the device).
An Observatory debugger and profiler on iPhone 12 Pro Max is available at: http://127.0.0.1:49593/7z6obr4MfsI=/

# 继续输入r，就进行热加载了
```

4、在模拟器中打开myapp的app就可以了

<----------👆创建项目就完成了👆-------->

### 打包build

#### 1、创建app签名

如果您有现有keystore，请跳至下一步。如果没有，请通过在运行以下命令来创建一个：` keytool -genkey -v -keystore ~/key.jks -keyalg RSA -keysize 2048 -validity 10000 -alias key`

🌟注意：` keytool是java jdk的一部分，也没有设置全局变量，所以路径是不对，它是作为Android Studio一部分安装的，找到对应路径，

设置私人密码123456

```
/Applications/Android Studio.app/Contents/jre/jdk/Contents/Home/jre/bin/keytool -genkey -v -keystore ~/key.jks -keyalg RSA -keysize 2048 -validity 10000 -alias key
```

✳️注意：还是报路径不存在

```
文件名中间有空格的，需要加个\  转义一下
```

根据提示填写密钥信息，会保存在~/key.jks这个文件下

#### 2、引用应用程序中的keystore

创建一个名为`<app dir>/android/key.properties`的文件，其中包含对密钥库的引用

```bash
storePassword=<password from previous step>
keyPassword=<password from previous step>
keyAlias=key
storeFile=<location of the key store file, e.g. /Users/<user name>/key.jks>
```

Eg:

```bash
storePassword=123456
keyPassword=123456
keyAlias=key
storeFile=/Users/fuguoyu/key.jks
```

#### 3、在gradle中配置签名

1）`<app dir>/android/app/build.gradle文件，在`android{`这一行前面,加入如下代码：

```
def keystorePropertiesFile = rootProject.file("key.properties")
def keystoreProperties = new Properties()
keystoreProperties.load(new FileInputStream(keystorePropertiesFile))

android {
```

2）替换

```bash
buildTypes {
    release {
        // TODO: Add your own signing config for the release build.
        // Signing with the debug keys for now, so `flutter run --release` works.
        signingConfig signingConfigs.debug
    }
}
```

为

```bash
signingConfigs {
    release {
        keyAlias keystoreProperties['keyAlias']
        keyPassword keystoreProperties['keyPassword']
        storeFile file(keystoreProperties['storeFile'])
        storePassword keystoreProperties['storePassword']
    }
}
buildTypes {
    release {
        signingConfig signingConfigs.release
    }
}
```

现在，应用的release版本将自动进行签名。

<----------👆打包配置[android]就完成了👆-------->

### 打包生成apk

直接在终端输入

```bash
flutter build apk
```



#### 安装真的没有问题么？咋没有网络请求？

今天flutter build apk打包了一个release.apk包，在真机上安装后网络数据都不显示，但是在模拟器上没问题，一切都正常！那这会是什么问题呢？问题所在，安卓开发中flutter应用没有网络权限

在这个文件里，android\app\src\profile\AndroidManifest.xml，manifest 里添加这段代码：

```
<manifest xmlns:android="http://schemas.android.com/apk/res/android"
....
  <uses-permission android:name="android.permission.READ_PHONE_STATE" />
  <uses-permission android:name="android.permission.INTERNET" />
  <uses-permission android:name="android.permission.ACCESS_NETWORK_STATE" />
  <uses-permission android:name="android.permission.ACCESS_WIFI_STATE" />
</manifest>
```

重新打包，虚拟机安装。还是没有数据。。。然后继续找，结果发现~~~

在路径android/src/main/AndroidManifest.xml，这里也有一个AndroidManifest.xml文件！跟之前的只不过是文件夹位置不同而已，同样在manifest标签下加入相同配置就行了，不要放到application里.

```
<manifest xmlns:android="http://schemas.android.com/apk/res/android"
....
  <uses-permission android:name="android.permission.READ_PHONE_STATE" />
  <uses-permission android:name="android.permission.INTERNET" />
  <uses-permission android:name="android.permission.ACCESS_NETWORK_STATE" />
  <uses-permission android:name="android.permission.ACCESS_WIFI_STATE" />
</manifest>
```

<----------👆打包apk就完成了👆-------->

完！