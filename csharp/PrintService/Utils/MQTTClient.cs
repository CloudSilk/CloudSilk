using MQTTnet.Client;
using MQTTnet;
using System;
using System.Threading;
using System.Threading.Tasks;
using System.Text.Json;
using MQTTnet.Server;
using Microsoft.Ajax.Utilities;

namespace PrintService.Utils
{
    internal static class ObjectExtensions
    {
        public static TObject DumpToConsole<TObject>(this TObject @object)
        {
            var output = "NULL";
            if (@object != null)
            {
                output = JsonSerializer.Serialize(@object, new JsonSerializerOptions
                {
                    WriteIndented = true
                });
            }

            Console.WriteLine($"[{@object?.GetType().Name}]:\r\n{output}");
            return @object;
        }
    }

    public class MQTTClient
    {
        public IMqttClient Client { get; set; }

        public MQTTClient(string host,int? port=null,Func<MqttApplicationMessageReceivedEventArgs,Task> handleReceived=null,string[] topics=null) {
            var mqttFactory = new MqttFactory();
            this.Client = mqttFactory.CreateMqttClient();
            var mqttClientOptions = new MqttClientOptionsBuilder().WithTcpServer(host,port).WithCredentials("admin","public").Build();
            if (handleReceived != null)
            {
                Client.ApplicationMessageReceivedAsync += handleReceived;
            }
            else
            {
                Client.ApplicationMessageReceivedAsync += HandleReceived;
            }
            this.Client.ConnectAsync(mqttClientOptions).Wait();
            if(topics != null)
            {
                this.Subscribe(topics).Wait();
            }
        }
        public Task HandleReceived(MqttApplicationMessageReceivedEventArgs e)
        {
            Console.WriteLine("Received application message.");
            e.DumpToConsole();

            return Task.CompletedTask;
        }

        public void Dispose()
        {
            Client?.DisconnectAsync(MqttClientDisconnectOptionsReason.NormalDisconnection);
            Client = null;
        }

        public Task<MqttClientSubscribeResult> Subscribe(string[] topics)
        {
            var mqttFactory = new MqttFactory();
            var builder = mqttFactory.CreateSubscribeOptionsBuilder();
            topics.ForEach(topic =>
            {
                builder.WithTopicFilter(topic);
            });
            return  Client.SubscribeAsync(builder.Build(), CancellationToken.None);
        }

        public Task<MqttClientPublishResult> Publish(string topic, string message)
        {
            var applicationMessage = new MqttApplicationMessageBuilder()
                .WithTopic(topic)
                .WithPayload(message)
                .Build();
            return Client.PublishAsync(applicationMessage, CancellationToken.None);
        }
    }
}