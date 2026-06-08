using UnityEngine;
using UnityEngine.UI;

public class SampleScene : MonoBehaviour
{
    public Button LoginButton;

    // Start is called once before the first execution of Update after the MonoBehaviour is created
    void Start()
    {
        LoginButton.onClick.AddListener(() =>
        {
            Debug.Log("haha");
        });
    }

    // Update is called once per frame
    void Update()
    {
        
    }
}