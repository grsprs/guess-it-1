FROM node:21-bookworm-slim

WORKDIR /app/

COPY package.json package-lock.json ./ 

RUN npm install

RUN apt-get update -y

RUN apt-get install golang python3 curl -y

RUN curl https://sh.rustup.rs/ -sSf | sh -s -- -y

COPY . /app/

RUN find /app -type f -name '*.sh' -exec sed -i 's/\r$//' {} + && \
    find /app -type f -name '*.go' -exec sed -i 's/\r$//' {} + && \
    find /app -type f -name '*.mod' -exec sed -i 's/\r$//' {} + && \
    chmod +x /app/script.sh /app/student/script.sh 2>/dev/null || true

CMD ["npm", "start"]
