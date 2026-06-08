
using Grpc.Core;
using UnityEngine;
using UnityEngine.UI;

public class SampleScene : MonoBehaviour
{
    public Button LoginButton;
    public Button ConnectButton;
    AuthService service;

    // Start is called once before the first execution of Update after the MonoBehaviour is created
    void Start()
    {
        LoginButton.onClick.AddListener(Login);
        ConnectButton.onClick.AddListener(Connect);
    }


    void Connect()
    {
        service ??= new();
        service.Connect();

    }

    void Login()
    {
        service.Login();
    }
}