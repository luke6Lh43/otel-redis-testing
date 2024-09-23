<?php

require 'vendor/autoload.php';

use Slim\Factory\AppFactory;
use Predis\Client as RedisClient;
use Psr\Http\Message\ResponseInterface as Response;
use Psr\Http\Message\ServerRequestInterface as Request;
use DI\ContainerBuilder;

// Create Container using PHP-DI
$containerBuilder = new ContainerBuilder();
$containerBuilder->addDefinitions([
    'redis' => function () {
        return new RedisClient([
            'scheme' => 'tcp',
            'host'   => '127.0.0.1',
            'port'   => 6379,
        ]);
    },
]);
$container = $containerBuilder->build();

// Create App with Container
AppFactory::setContainer($container);
$app = AppFactory::create();

// Define route to save and retrieve a simple text message
$app->get('/data', function (Request $request, Response $response, $args) {
    $redis = $this->get('redis');

    $cacheKey = 'simple_text_message';
    $cachedMessage = $redis->get($cacheKey);

    if ($cachedMessage) {
        $message = $cachedMessage;
    } else {
        $message = 'Hello, this is a simple text message!';
        $redis->setex($cacheKey, 3600, $message); // Cache for 1 hour
    }

    $response->getBody()->write($message);
    return $response->withHeader('Content-Type', 'text/plain');
});

// Run App
$app->run();
