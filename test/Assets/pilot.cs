using System.Collections;
using System.Collections.Generic;
using UnityEngine;

public class pilot : MonoBehaviour
{
    public float thrust;
    public float inertia;
    public float current_inertia;

    public float h_inertia;
    public float h_current_inertia;
    public Rigidbody2D rb;
    public GameObject mt;
    public GameObject rt;
    public GameObject lt;


    // Start is called before the first frame update
    void Start()
    {    }

    void HorizontalMove()
    {
        float translation = Input.GetAxis("Horizontal");
        if (translation != 0){

            h_current_inertia = Mathf.Min(h_inertia, h_current_inertia + translation);
            h_current_inertia = Mathf.Max(-h_inertia, h_current_inertia + translation);
            //transform.Translate(new Vector3(0, thrust * current_inertia / inertia, 0)) ;
            if(translation > 0) lt.SetActive(true);
            else rt.SetActive(true);
        }
        else
        {
            h_current_inertia *= 999/1000 ;
            lt.SetActive(false);
            rt.SetActive(false);
        }
        rb.AddForce(new Vector3((thrust/2) * h_current_inertia / h_inertia, 0, 0));
    }

    // Update is called once per frame <>
    void Update()
    {
        // Get the horizontal and vertical axis.
        // By default they are mapped to the arrow keys.
        // The value is in the range -1 to 1
        float translation = Input.GetAxis("Vertical");
        if (translation > 0)
        {
            current_inertia = Mathf.Min(inertia, current_inertia + translation);
            mt.SetActive(true);
            //transform.Translate(new Vector3(0, thrust * current_inertia / inertia, 0)) ;

            translation = thrust;
        }
        else
        {
            current_inertia = Mathf.Max(0, current_inertia - 0.4f) ;
            mt.SetActive(false);
            translation = 1f;
        }
        rb.AddForce(new Vector3(0, translation * current_inertia / inertia, 0));

        HorizontalMove();
    }

    private void OnTriggerEnter2D(Collider2D collision)
    {
        Debug.Log("Booom 2");
        transform.SetPositionAndRotation(new Vector3(-4, 0, 0), new Quaternion());
        h_current_inertia = 0;
        current_inertia = 0;
    }
}
