FROM openjdk:17-jdk-bullseye

# https://developer.android.com/ndk/downloads?hl=zh-cn
ENV NDK_LTS_VERSION 27.2.12479018
# https://hub.docker.com/layers/cimg/android/2025.04.1-ndk/images/sha256-3a4de864bc5e9b943e18e4d33bb8553d4f82fde6c21c364986ad556ccc8681d4 41
ENV SDK_TOOLS_VERSION 13114758
ENV ANDROID_PLATFORM_VERSION 31

### Adapted from CircleCI Android docker image ###
# https://github.com/CircleCI-Public/cimg-android/blob/main/2022.03/Dockerfile
# https://github.com/CircleCI-Public/cimg-android/blob/7af5c262821e4bf68c19cdb0ed34283165684e11/2022.03/Dockerfile
# NDK part:
# https://github.com/CircleCI-Public/cimg-android/blob/main/2022.03/ndk/Dockerfile
# https://github.com/CircleCI-Public/cimg-android/blob/7af5c262821e4bf68c19cdb0ed34283165684e11/2022.03/ndk/Dockerfile

ENV ANDROID_HOME "/home/circleci/android-sdk"
ENV ANDROID_SDK_ROOT $ANDROID_HOME
ENV CMDLINE_TOOLS_ROOT "${ANDROID_HOME}/cmdline-tools/latest/bin"
ENV ADB_INSTALL_TIMEOUT 120
ENV PATH "${ANDROID_HOME}/emulator:${ANDROID_HOME}/cmdline-tools/latest/bin:${ANDROID_HOME}/tools:${ANDROID_HOME}/tools/bin:${ANDROID_HOME}/platform-tools:${ANDROID_HOME}/platform-tools/bin:${PATH}"
ENV ANDROID_NDK_HOME "/home/circleci/android-sdk/ndk/${NDK_LTS_VERSION}"
ENV ANDROID_NDK_ROOT "${ANDROID_NDK_HOME}"

RUN set -eux && \
    mkdir -p ${ANDROID_HOME}/cmdline-tools && \
	mkdir ${ANDROID_HOME}/platforms && \
	mkdir ${ANDROID_HOME}/ndk && \
	wget -O /tmp/cmdline-tools.zip -t 5 --no-verbose \
        "https://dl.google.com/android/repository/commandlinetools-linux-${SDK_TOOLS_VERSION}_latest.zip" && \
	unzip -q /tmp/cmdline-tools.zip -d ${ANDROID_HOME}/cmdline-tools && \
	rm /tmp/cmdline-tools.zip && \
	mv ${ANDROID_HOME}/cmdline-tools/cmdline-tools ${ANDROID_HOME}/cmdline-tools/latest && \

    # Use sdkmanager to install further tools

	echo y | ${CMDLINE_TOOLS_ROOT}/sdkmanager "build-tools;${ANDROID_PLATFORM_VERSION}.0.0" && \       
    echo y | ${CMDLINE_TOOLS_ROOT}/sdkmanager "platforms;android-${ANDROID_PLATFORM_VERSION}" && \
    echo y | ${CMDLINE_TOOLS_ROOT}/sdkmanager "ndk;${NDK_LTS_VERSION}"
    #echo y | ${CMDLINE_TOOLS_ROOT}/sdkmanager "platform-tools" && \
    #echo y | ${CMDLINE_TOOLS_ROOT}/sdkmanager "tools" && \
    #echo y | ${CMDLINE_TOOLS_ROOT}/sdkmanager "cmake;3.10.2.4988404" && \
	#echo y | ${CMDLINE_TOOLS_ROOT}/sdkmanager "cmake;3.18.1" && \

### End of adaption ###


# Go parts adapted from CircleCI Go docker image
# https://github.com/CircleCI-Public/cimg-go/blob/main/1.17/Dockerfile

ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH
ENV GO_VERSION 1.23.9

# Install packages needed for CGO
RUN apt-get update && apt-get install -y --no-install-recommends \
		g++ \
        make \
        git \
		libc6-dev && \
	rm -rf /var/lib/apt/lists/* && \
	curl -sSL "https://golang.org/dl/go${GO_VERSION}.linux-amd64.tar.gz" | tar -xz -C /usr/local/ && \
	mkdir -p $GOPATH/bin
    
    # Install and setup gomobile now
    # This is not from CircleCI

RUN set -eux; \
    go install "golang.org/x/mobile/cmd/gomobile@latest"; \
    gomobile init; \
    mkdir /module

VOLUME "/module"
WORKDIR "/module"

ENTRYPOINT ["gomobile"]

