FROM python:alpine

WORKDIR /app
COPY ["run.sh", "requirement.txt", "./"]

RUN pip install --no-cache-dir -r requirement.txt

COPY . .

CMD ["./run.sh"]