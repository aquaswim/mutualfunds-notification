FROM python:alpine

RUN apk add --no-cache tzdata

WORKDIR /app
COPY ["run.sh", "requirements.txt", "./"]

RUN pip install --no-cache-dir -r requirements.txt

COPY . .

CMD ["./run.sh"]