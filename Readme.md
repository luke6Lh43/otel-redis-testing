# OpenTelemetry Redis Attribute Verification

This project aims to verify the attributes sent to the OpenTelemetry Collector for Redis across six of the most popular programming languages: Java, .NET, Python, PHP, Node.js, and Go. The goal is to ensure alignment with the latest versions of OpenTelemetry instrumentation libraries and the OpenTelemetry Semantic Conventions for Redis.

## Objective

The primary objective of this project is to validate that the attributes generated by the OpenTelemetry instrumentation libraries for Redis are consistent with the [OpenTelemetry Semantic Conventions for Redis](https://opentelemetry.io/docs/specs/semconv/database/redis/).

## Languages Covered

- Java
- .NET
- Python
- PHP
- Node.js
- Go

## How to Use

1. Clone this repository.
2. Deploy Redis and Jaeger:

```
docker run -d --name jaeger -e COLLECTOR_OTLP_ENABLED=true -p 4317:4317 -p 4318:4318 -p 16686:16686 jaegertracing/all-in-one:latest
docker run --name redis -p 6379:6379 -d redis
```

3. Follow the setup instructions for each language-specific directory:

### .NET
*Tested with dotnet 8.0.104*

```
cd dotnet-redis
dotnet build

OpenTelemetry Zero-code Instrumentation:

curl -L -O https://github.com/open-telemetry/opentelemetry-dotnet-instrumentation/releases/latest/download/otel-dotnet-auto-install.sh
chmod +x $HOME/.otel-dotnet-auto/instrument.sh
. $HOME/.otel-dotnet-auto/instrument.sh

export OTEL_SERVICE_NAME=dotnet-redis-testing
export OTEL_EXPORTER_OTLP_PROTOCOL=http/protobuf
export OTEL_EXPORTER_OTLP_ENDPOINT=http://localhost:4318
export OTEL_TRACES_EXPORTER=console,otlp
export OTEL_PROPAGATORS=baggage,tracecontext
export OTEL_DOTNET_AUTO_TRACES_CONSOLE_EXPORTER_ENABLED=true

Run the .NET application:

dotnet run
```

### GO
*Tested with go1.23.1*

```
cd go-redis
go mod tidy
go build
go run main.go
```

### JAVA
*Tested with openjdk 22.0.2, Apache Maven 3.9.8*

```
cd java-redis
mvn clean install

wget https://github.com/open-telemetry/opentelemetry-java-instrumentation/releases/download/v2.8.0/opentelemetry-javaagent.jar

java -javaagent:./opentelemetry-javaagent.jar -Dotel.exporter.otlp.endpoint=http://localhost:4318 -Dotel.exporter.otlp.protocol=http/protobuf  -Dotel.metrics.exporter=none -Dotel.logs.exporter=none -Dotel.resource.attributes=service.name=java-redis-testing -Dotel.traces.exporter=otlp,console -jar target/java-redis-1.0-SNAPSHOT.jar
```

### Node.js
*Tested with Node.js 22.9.0, NPM 10.8.3*

```
cd nodejs-redis
npm install
node -r ./tracing.js app.js

curl http://localhost:3000
```

### PHP
*Tested with PHP 8.3.11, Composer 2.7.9*

```
cd php-redis

#install opentelemetry extension
pecl install opentelemetry

#please double-check what is extension directory and update php.ini file if needed
php -i | grep extension_dir
cat php.ini | grep opentelemetry

composer install

export OTEL_PHP_AUTOLOAD_ENABLED=true
export OTEL_SERVICE_NAME=php-redis-testing

#For some reason OTEL Traces Exporter for PHP only support one exporter, so env variable must be either "console" or "otlp"
export OTEL_TRACES_EXPORTER=console

export OTEL_EXPORTER_OTLP_PROTOCOL=http/protobuf
export OTEL_EXPORTER_OTLP_ENDPOINT=http://localhost:4318
export OTEL_PROPAGATORS=baggage,tracecontext

php -S localhost:8080 -t .

curl http://localhost:8080/data
```

### Python
*Tested with Python 3.12.5*

```
cd python-redis
python -m venv venv
source venv/bin/activate
pip install -r requirements.txt
opentelemetry-bootstrap -a install

opentelemetry-instrument --traces_exporter console,otlp --exporter_otlp_endpoint http://0.0.0.0:4317 --service_name python-redis-testing python app.py
```