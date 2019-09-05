using System.Collections;
using System.Collections.Generic;
using UnityEngine;

public class thruter : MonoBehaviour
{
    // Start is called before the first frame update
    void Start()
    {
        
    }

    // Update is called once per frame
    void Update()
    {
        float translation = Input.GetAxis("Vertical");
        if (translation > 0)
        {
            Debug.Log("hsfhe");
        }
    }
}
