const { NodeSDK } = require('@opentelemetry/sdk-node');
const { getNodeAutoInstrumentations } = require('@opentelemetry/auto-instrumentations-node');
const { OTLPTraceExporter } = require('@opentelemetry/exporter-trace-otlp-http');
const { ConsoleSpanExporter, BatchSpanProcessor } = require('@opentelemetry/sdk-trace-base');
const { diag, DiagConsoleLogger, DiagLogLevel } = require('@opentelemetry/api');
const { Resource } = require('@opentelemetry/resources');
const { SemanticResourceAttributes } = require('@opentelemetry/semantic-conventions');

// For troubleshooting, set the log level to DiagLogLevel.DEBUG
diag.setLogger(new DiagConsoleLogger(), DiagLogLevel.INFO);

const otlpTraceExporter = new OTLPTraceExporter({
  url: 'http://localhost:4318/v1/traces', // Default endpoint for OTLP HTTP
});

const consoleTraceExporter = new ConsoleSpanExporter();

const sdk = new NodeSDK({
  spanProcessors: [
    new BatchSpanProcessor(otlpTraceExporter),
    new BatchSpanProcessor(consoleTraceExporter),
  ],
  instrumentations: [
    getNodeAutoInstrumentations({
      '@opentelemetry/instrumentation-fs': {
        enabled: false, // Disable fs instrumentation
      },
    }),
  ],
  resource: new Resource({
    [SemanticResourceAttributes.SERVICE_NAME]: 'nodejs-redis-testing', // Custom service name
  }),
});

(async () => {
  try {
    await sdk.start();
    console.log('Tracing initialized');
  } catch (error) {
    console.error('Error initializing tracing', error);
  }
})();

// Gracefully shut down the SDK on process exit
process.on('SIGTERM', async () => {
  try {
    await sdk.shutdown();
    console.log('Tracing terminated');
  } catch (error) {
    console.error('Error terminating tracing', error);
  } finally {
    process.exit(0);
  }
});

