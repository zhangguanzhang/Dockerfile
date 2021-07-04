## usage

```
docker run --name apktool --rm -tid -v $PWD:/opt zhangguanzhang/apktool
```

### 反编译

```
apktool d xx.apk
```

```
cd xx
vi res/xml/network_security_config.xml
<network-security-config>
    <base-config cleartextTrafficPermitted="true">
        <trust-anchors>
            <certificates src="system" overridePins="true" />
            <certificates src="user" overridePins="true" />
        </trust-anchors>
    </base-config>
</network-security-config>
```

解包开的根目录下`AndroidManifest.xml`里的<application>标签中确保有`android:networkSecurityConfig="@xml/network_security_config"`

```
$ grep -n android:networkSecurityConfig AndroidManifest.xml
74:    <application android:allowBackup="false" android:appComponentFactory="androidx.core.app.CoreComponentFactory" android:icon="@mipmap/ic_launcher" android:label="AcFun" android:largeHeap="true" android:logo="@drawable/icon_actionbar" android:name="tv.acfun.core.application.AcFunApplication" android:networkSecurityConfig="@xml/network_security_config" android:roundIcon="@mipmap/ic_launcher_round" android:supportsRtl="true" android:theme="@style/AppTheme.Light">
```

### 重新打包

```
apktool b xx
ls -l xx/dist/

# 暂时不能用貌似
keytool -genkey -alias aeo_android.keystore -keyalg RSA -validity 20000 -keystore aeo_android.keystore

jarsigner -verbose -keystore aeo_android.keystore -signedjar new_xx.apk xx.apk aeo_android.keystore

# 来源 https://github.com/as0ler/Android-Tools/tree/master/Autosign/Auto-Sign
java -jar signapk.jar testkey.x509.pem testkey.pk8 xx.apk new_xx.apk
```


## justTrustMe

```

git clone https://github.com/JunGe-Y/JustTrustMePP.git
cd JustTrustMePP

APK_PATH="bin/JustTrustMe.apk"
mkdir bin
./gradlew assembleRelease && cp app/build/outputs/apk/app-release-unsigned.apk $APK_PATH && signapk $APK_PATH

# --------
git clone https://github.com/Fuzion24/JustTrustMe.git
cd JustTrustMe

docker run -v $PWD:/app --workdir /app -ti --entrypoint bash --user root  febririzki46/gradlew-android

apt-get update
apt-get install -y signapk



https://developer.android.com/studio/#downloads

wget https://dl.google.com/android/repository/commandlinetools-linux-7302050_latest.zip

# unpack archive
unzip sdk-tools-linux-4333796.zip

```

## 参考

- https://segmentfault.com/a/1190000023087971
- https://developer.android.com/training/articles/security-config