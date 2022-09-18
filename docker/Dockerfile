FROM gitpod/workspace-full

#RUN sudo apt-get update && \
#    sudo apt-get install gettext libncurses5 libxkbcommon0 libtinfo5 -y

USER gitpod

RUN brew install httpie && \
    #brew install bat && \
    brew install exa && \
    #brew install hey && \
    brew install redis

# ------------------------------------
# Install TinyGo
# ------------------------------------
ARG TINYGO_VERSION="0.24.0"
RUN wget https://github.com/tinygo-org/tinygo/releases/download/v${TINYGO_VERSION}/tinygo_${TINYGO_VERSION}_amd64.deb
RUN sudo dpkg -i tinygo_${TINYGO_VERSION}_amd64.deb
RUN rm tinygo_${TINYGO_VERSION}_amd64.deb

# ------------------------------------
# Install NATS server
# ------------------------------------

ARG NATS_VERSION="2.9.0"
ARG NATS_OS="linux-amd64"
RUN curl -L https://github.com/nats-io/nats-server/releases/download/v${NATS_VERSION}/nats-server-v${NATS_VERSION}-${NATS_OS}.zip -o nats-server.zip && \
    unzip nats-server.zip -d nats-server && \
    sudo cp nats-server/nats-server-v${NATS_VERSION}-${NATS_OS}/nats-server /usr/bin && \
    rm nats-server.zip && \
    rm -rf nats-server
