

using System;
using System.Threading.Tasks;
using Cysharp.Net.Http;
using Dakg;
using Grpc.Net.Client;

public class AuthService
{

    GrpcChannel _channel;
    Dakg.GameService.GameServiceClient _client;

    public void Connect()
    {
        _channel = GrpcChannel.ForAddress("http://127.0.0.1:50051", new GrpcChannelOptions
        {
            HttpHandler = new YetAnotherHttpHandler { Http2Only = true },
            DisposeHttpClient = true
        });
        _client = new(_channel);
    }

    public async Task<LoginResponse> Login()
    {
        LoginRequest req = new()
        {
            Publickey = "haha"
        };
        var result = await _client.LoginAsync(req);
        return result; ;
    }
}