### DOTNET

```csharp
Retrieved value: sampleValue
Activity.TraceId:            5e54c59520df46ad76daf2a83dea4b7f
Activity.SpanId:             3c2faed9c113d090
Activity.TraceFlags:         Recorded
Activity.ActivitySourceName: OpenTelemetry.Instrumentation.StackExchangeRedis
Activity.ActivitySourceVersion: 1.0.0-rc9.14
Activity.DisplayName:        SET
Activity.Kind:               Client
Activity.StartTime:          2024-09-23T16:42:05.7827280Z
Activity.Duration:           00:00:00.0044677
Activity.Tags:
    db.system: redis
    db.redis.flags: DemandMaster
    db.statement: SET
    net.peer.name: localhost
    net.peer.port: 6379
    db.redis.database_index: 0
```    
### GO

```go
"Attributes": [
        {
                "Key": "server.address",
                "Value": {
                        "Type": "STRING",
                        "Value": "localhost"
                }
        },
        {
                "Key": "server.port",
                "Value": {
                        "Type": "INT64",
                        "Value": 6379
                }
        },
        {
                "Key": "db.system",
                "Value": {
                        "Type": "STRING",
                        "Value": "redis"
                }
        },
        {
                "Key": "db.connection_string",
                "Value": {
                        "Type": "STRING",
                        "Value": "redis://localhost:6379"
                }
        },
        {
                "Key": "code.function",
                "Value": {
                        "Type": "STRING",
                        "Value": "main.main"
                }
        },
        {
                "Key": "code.filepath",
                "Value": {
                        "Type": "STRING",
                        "Value": "/Users/luke6Lh43/developer/opentelemetry/otel-redis-testing/go-redis/main.go"
                }
        },
        {
                "Key": "code.lineno",
                "Value": {
                        "Type": "INT64",
                        "Value": 198
                }
        },
        {
                "Key": "db.statement",
                "Value": {
                        "Type": "STRING",
                        "Value": "get id1234"
                }
        }
]
```

## JAVA

```java

21 78b6d9336e1ee84e CLIENT [tracer: io.opentelemetry.jedis-4.0:2.8.0-alpha] AttributesMap{data={db.operation=GET, network.peer.port=6379, thread.name=main, thread.id=1, network.type=ipv6, network.peer.address=0:0:0:0:0:0:0:1, db.statement=GET mykey, db.system=redis}, capacity=128, totalAddedValues=8}
Retrieved value: myvalue
```

## NODEJS

```node

  instrumentationScope: {
    name: '@opentelemetry/instrumentation-redis-4',
    version: '0.42.0',
    schemaUrl: undefined
  },
  traceId: '2ae19e231fa0e53bf5c4f381960a08f8',
  parentId: '93af18f4a8887053',
  traceState: undefined,
  name: 'redis-GET',
  id: '9593cfff9135160d',
  kind: 2,
  timestamp: 1727193579009000,
  duration: 1111.802,
  attributes: {
    'db.system': 'redis',
    'net.peer.name': 'localhost',
    'net.peer.port': 6379,
    'db.connection_string': 'redis://localhost:6379',
    'db.statement': 'GET myKey'
  },
  status: { code: 0 },
  events: [],
  links: []
}
```

## PHP

```php

"attributes": {
    "code.function": "__construct",
    "code.namespace": "Predis\\Client",
    "code.filepath": "\/Users\/luke6Lh43\/developer\/opentelemetry\/otel-redis-testing\/php-redis\/vendor\/predis\/predis\/src\/Client.php",
    "code.lineno": 71,
    "server.address": "127.0.0.1",
    "network.transport": "tcp",
    "server.port": 6379,
    "db.system": "redis"
}
```

## PYTHON

```python
{
    "name": "GET",
    "context": {
        "trace_id": "0xeae050be5f998589444a7f56e0a3f35a",
        "span_id": "0x5a90622d4fd33ff9",
        "trace_state": "[]"
    },
    "kind": "SpanKind.CLIENT",
    "parent_id": null,
    "start_time": "2024-09-23T17:18:38.663355Z",
    "end_time": "2024-09-23T17:18:38.663780Z",
    "status": {
        "status_code": "UNSET"
    },
    "attributes": {
        "db.statement": "GET ?",
        "db.system": "redis",
        "db.redis.database_index": 0,
        "net.peer.name": "localhost",
        "net.peer.port": 6379,
        "net.transport": "ip_tcp",
        "db.redis.args_length": 2
    },
    "events": [],
    "links": [],
    "resource": {
        "attributes": {
            "telemetry.sdk.language": "python",
            "telemetry.sdk.name": "opentelemetry",
            "telemetry.sdk.version": "1.27.0",
            "service.name": "python-redis-testing",
            "telemetry.auto.version": "0.48b0"
        },
        "schema_url": ""
    }
}
```