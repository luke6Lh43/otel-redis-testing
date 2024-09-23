using System;
using StackExchange.Redis;

class Program
{
    private static void Main(string[] args)
    {
        var redis = ConnectionMultiplexer.Connect("localhost:6379");
        var db = redis.GetDatabase();

        var key = "sampleKey";
        var value = "sampleValue";

        // Save record
        db.StringSet(key, value);

        // Get record
        var retrievedValue = db.StringGet(key);

        // Print record
        Console.WriteLine($"Retrieved value: {retrievedValue}");
    }
}