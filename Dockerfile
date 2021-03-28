FROM golang:1.16.2-buster
RUN mkdir /app
WORKDIR /app
# Depedencies:  https://ebiten.org/documents/install.html
RUN apt update && \
apt install -y \
xvfb \
libc6-dev \
libglu1-mesa-dev \
libgl1-mesa-dev \
libxcursor-dev \
libxi-dev \
libxinerama-dev \
libxrandr-dev \
libxxf86vm-dev \
libasound2-dev \
pkg-config \
libalut-dev

ENV DISPLAY :99

COPY . .
RUN ./build.sh
EXPOSE 8080

CMD ["/app/run-web.sh"]