

using System.Threading.Tasks;
using Dakg;
using Grpc.Net.Client;

public class AuthService
{

    Dakg.GameService.GameServiceClient _client;

    public void Connect()
    {
        _client = new(GrpcChannel.ForAddress("https://127.0.0.1:50051"));
    }

    public async Task<LoginResponse> Login()
    {
        LoginRequest req = new();
        var result = await _client.LoginAsync(req);
        return result; ;
    }

}